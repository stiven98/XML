import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BlockUserDTO } from '../model/BlockUser';
import { DeletePost } from '../model/DeletePost';
import { ReportReq } from '../model/ReportReq';
import { AuthService } from '../service/authorization/auth.service';
import { PostsService } from '../service/posts.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-reported-posts',
  templateUrl: './reported-posts.component.html',
  styleUrls: ['./reported-posts.component.css'],
})
export class ReportedPostsComponent implements OnInit {
  public posts: any[] = [];
  public userData: Map<string, any> = new Map<string, any>();
  public reportReq: ReportReq = new ReportReq();
  public blockUser: BlockUserDTO = new BlockUserDTO();
  public deletePost: DeletePost = new DeletePost();
  constructor(
    public authService: AuthService,
    public userService: UserService,
    public postsService: PostsService,
    public router: Router
  ) {}

  ngOnInit(): void {
    this.initData();
  }
  initData = () => {
    this.postsService.getReportedPosts().subscribe((res) => {
      this.posts = res as any[];
      for (let p of this.posts) {
        if (!this.userData.get(p.userid)) {
          this.userService.getUsersById(p.userid).subscribe((res) => {
            this.userData.set(p.userid, res);
          });
        }
      }
    });
  };

  convertTimeToText = (timeString: string) => {
    let time = new Date(timeString);
    let now = new Date();
    let diff = now.valueOf() - time.valueOf();
    if (diff >= 86400000) {
      return Math.ceil(diff / 86400000) + ' days ago';
    }
    if (diff >= 3600000) {
      return Math.ceil(diff / 3600000) + ' hours ago';
    }
    if (diff >= 60000) {
      return Math.ceil(diff / 60000) + ' minutes ago';
    } else {
      return Math.ceil(diff / 1000) + ' seconds ago';
    }
  };

  likeClick = (post: any) => {
    let postId = post.id;
    let ownerId = post.userid;
    let id = localStorage.getItem('id');
    if (id) {
      this.postsService
        .likePost({ userid: id, postid: postId, ownerid: ownerId })
        .subscribe((res) => {
          this.ngOnInit();
        });
    } else {
      alert('Morate biti ulogovani da bi lajkovali objavu');
    }
  };

  dislikeClick = (post: any) => {
    let postId = post.id;
    let ownerId = post.userid;
    let id = localStorage.getItem('id');
    if (id) {
      this.postsService
        .dislikePost({ userid: id, postid: postId, ownerid: ownerId })
        .subscribe((res) => {
          this.ngOnInit();
        });
    } else {
      alert('Morate biti ulogovani da bi dislajkovali objavu');
    }
  };

  reportPost = (postId: string, ownerId: string) => {
    this.reportReq.PostId = postId;
    this.reportReq.UserId = localStorage.getItem('id') as string;
    this.reportReq.OwnerId = ownerId;
    this.postsService.reportPost(this.reportReq).subscribe((res) => res);
  };

  isImage = (name: string): boolean => {
    let imgFormats = ['jpg', 'jpeg', 'gif', 'png', 'tiff', 'bmp'];
    let flag = false;
    for (let i = 0; i < imgFormats.length; i++) {
      if (name.includes(imgFormats[i])) {
        flag = true;
        break;
      }
    }
    return flag;
  };

  onDelete = (id: string, postId: string) => {
    this.deletePost.ownerId = postId;
    this.deletePost.postId = id;
    this.postsService.deletePost(this.deletePost).subscribe((res) => {
      if (res) {
        alert('Uspešno ste obrisali objavu');
      } else {
        alert('Došlo je do greške');
      }
      this.initData();
    });
  };
  onBlock = (id: string, postId: string) => {
    this.deletePost.ownerId = id;
    this.deletePost.postId = postId;
    this.blockUser.userId = id;
    this.userService.blockUser(this.blockUser).subscribe((res) => {
      if (res) {
        this.postsService
          .deletePost(this.deletePost)
          .subscribe((result) => result);
        alert('Uspešno ste blokirali korisnika');
      } else {
        alert('Došlo je do greške');
      }
      this.initData();
    });
  };
}
