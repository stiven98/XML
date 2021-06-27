import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ReportReq } from '../model/ReportReq';
import { AuthService } from '../service/authorization/auth.service';
import { PostsService } from '../service/posts.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-single-post',
  templateUrl: './single-post.component.html',
  styleUrls: ['./single-post.component.css']
})
export class SinglePostComponent implements OnInit {

  public post: any;
  public userData: Map<string, any> = new Map<string, any>();
  public commentData: Map<string, string> = new Map<string, string>();
  public reportReq: ReportReq = new ReportReq();
  public currentUser: any;
  public commentText: string = '';
  public userid : string = '';
  public postid : string = '';

  constructor(public postsService: PostsService,
    public userService: UserService,
    public authService: AuthService,
    public router: Router,
    public route: ActivatedRoute) {
  }

  convertTimeToText = (timeString: string) => {
    let time = new Date(timeString);
    let now = new Date();
    let diff = now.valueOf() - time.valueOf();
    if (diff >= 86400000) {
      return Math.floor(diff / 86400000) + ' days ago';
    }
    if (diff >= 3600000) {
      return Math.floor(diff / 3600000) + ' hours ago';
    }
    if (diff >= 60000) {
      return Math.floor(diff / 60000) + ' minutes ago';
    } else {
      return Math.floor(diff / 1000) + ' seconds ago';
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
          this.initData();
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
          this.initData();
        });
    } else {
      alert('Morate biti ulogovani da bi dislajkovali objavu');
    }
  };

  reportPost = (postId: string, ownerId: string) => {
    this.reportReq.PostId = postId;
    this.reportReq.UserId = localStorage.getItem('id') as string;
    this.reportReq.OwnerId = ownerId;
    this.postsService.reportPost(this.reportReq).subscribe((res) => {this.initData()});
  };

  ngOnInit(): void {
    this.postid = this.route.snapshot.params[`postid`];
    this.userid = this.route.snapshot.params[`userid`];
    this.initData();
  }

  initData = () => {
    let id = localStorage.getItem('id');
    this.postsService.getByIds(this.userid, this.postid).subscribe((res) => {
      this.post = res as any;
      this.userService.getUsersById(id).subscribe(response => {
        this.currentUser = response;
      });
      if (this.post) {
          if (!this.userData.get(this.post.userid)) {
            this.userService.getUsersById(this.post.userid).subscribe(res => {
              this.userData.set(this.post.userid, res);
            });
          }
      this.fetchCommentData();
      }
    });
  }

  fetchCommentData = () => {
    if (this.post) {
        for (let c of this.post.comments) {
          if (!this.userData.get(c.userid)) {
            this.userService.getUsersById(c.userid).subscribe(res => {
              this.userData.set(c.userid, res);
            });
          }
        }
    }
  }

  getCurrentUserImage = (): string => {
    return this.currentUser.system_user.picturePath;
  }

  getUserImage = (comment: any): string => {
    var user = this.userData.get(comment.userid);
    return user.system_user.picturePath;
  }

  getUsername = (id: string): string => {
    var user = this.userData.get(id);
    if (user)
      return user.system_user.username;
    else
      return '';
  }

  leaveComment = (post: any) => {
    let id = localStorage.getItem('id');
    if (id) {
      var commentValue = this.commentData.get(post.id) as string;
      this.postsService.leaveComment({ ownerid: post.userid, postid: post.id, userid: id, comment: commentValue }).subscribe(res => {
        this.initData();
      });
    } else {
      alert('Morate biti ulogovani da bi komentarisali objavu')
    }
  }

  changeCommentText = (event: Event, id: string) => {
    const e = event.target as HTMLInputElement;
    let value = e.value;
    this.commentData.set(id, value);
  }

  onUsernameClick = (event: Event, id: string) => {
    event.preventDefault();
    var user = this.userData.get(id);
    if (user)
      this.router.navigate(['/profile', id]);
  }


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
  }

}
