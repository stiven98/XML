import { Component, OnInit } from '@angular/core';
import { ReportReq } from 'src/app/model/ReportReq';
import { AuthService } from 'src/app/service/authorization/auth.service';
import { PostsService } from 'src/app/service/posts.service';
import { UserService } from 'src/app/service/user.service';


@Component({
  selector: 'app-liked-disliked-posts',
  templateUrl: './liked-disliked-posts.component.html',
  styleUrls: ['./liked-disliked-posts.component.css'],
})
export class LikedDislikedPostsComponent implements OnInit {
  public posts: any[] = [];
  public userData: Map<string, any> = new Map<string, any>();
  public reportReq: ReportReq = new ReportReq();

  constructor(
    public authService: AuthService,
    public userService: UserService,
    public postsService: PostsService
  ) {}

  ngOnInit(): void {
    let id = localStorage.getItem('id') as string;
    if (window.location.href.includes('disliked')) {
      this.postsService.getDislikedPosts(id).subscribe((res) => {
        this.posts = res as any[];
        for (let p of this.posts) {
          if (!this.userData.get(p.userid)) {
            this.userService.getUsersById(p.userid).subscribe(res => {
              this.userData.set(p.userid, res);
            });
          }
        }
      });
    } else {
      this.postsService.getLikedPosts(id).subscribe((res) => {
        this.posts = res as any[];
        for (let p of this.posts) {
          if (!this.userData.get(p.userid)) {
            this.userService.getUsersById(p.userid).subscribe(res => {
              this.userData.set(p.userid, res);
            });
          }
        }
      })
    }
  }

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
    for(let i = 0; i < imgFormats.length; i++){
      if(name.includes(imgFormats[i])){
        flag = true;
        break;
      }
    }
    return flag;
  }
}
