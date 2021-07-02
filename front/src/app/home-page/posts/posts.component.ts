import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ReportReq } from 'src/app/model/ReportReq';
import { AuthService } from 'src/app/service/authorization/auth.service';
import { FollowService } from 'src/app/service/follow.service';
import { NotifyService } from 'src/app/service/notify.service';
import { PostsService } from 'src/app/service/posts.service';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {
  @Input() tagFilter: boolean = false;
  @Input() locationFilter: boolean = false;
  @Input() tagForSearch: string = '';
  @Input() locationForSearch: string = '';
  public posts: any[] = [];
  public postsToShow: any[] = [];
  public publicPosts: any[] = [];
  public userData: Map<string, any> = new Map<string, any>();
  public commentData: Map<string, string> = new Map<string, string>();
  public reportReq: ReportReq = new ReportReq();
  public showPublic: boolean = false;
  public currentUser: any;
  public commentText: string = '';

  constructor(public postsService: PostsService,
    public userService: UserService,
    public authService: AuthService,
    public router: Router,
    public notifyService: NotifyService,
    public followService: FollowService) {
    if (!this.authService.isLoggedIn) {
      this.showPublic = true;
    }
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
        .subscribe(async (res) => {
          let flag = await this.isNotificationNeededLikes(this.userData.get(ownerId), this.currentUser);
          if (flag) {
            this.notifyService.notifyUser(
              {
                "userId": id,
                "notify_user_id": ownerId,
                "type_of_notify": 'like',
                "notify_id": postId
              }
            ).subscribe(res => res);
          }
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
        .subscribe(async (res) => {
          let flag = await this.isNotificationNeededDislikes(this.userData.get(ownerId), this.currentUser);
          if (flag) {
            this.notifyService.notifyUser(
              {
                "userId": id,
                "notify_user_id": ownerId,
                "type_of_notify": 'dislike',
                "notify_id": postId
              }
            ).subscribe(res => res);
          }
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
    this.postsService.reportPost(this.reportReq).subscribe((res) => { this.initData() });
  };

  ngOnInit(): void {
    this.initData();
  }

  initData = () => {
    let id = localStorage.getItem('id');
    this.postsService.getPublicPosts(id as string).subscribe((res) => {
      this.publicPosts = res as any[];
      if (this.publicPosts) {
        for (let p of this.publicPosts) {
          if (!this.userData.get(p.userid)) {
            this.userService.getUsersById(p.userid).subscribe(res => {
              this.userData.set(p.userid, res);
            });
          }
        }
      }

      if (id) {
        this.postsService.getFeed(id).subscribe(res => {
          this.posts = res as any[];
          if (this.posts) {
            for (let p of this.posts) {
              if (!this.userData.get(p.userid)) {
                this.userService.getUsersById(p.userid).subscribe(res => {
                  this.userData.set(p.userid, res);
                });
              }
            }
          }

          this.userService.getUsersById(id).subscribe(response => {
            this.currentUser = response;
          });

          this.setPostsToDisplay();
        });
      }
      this.setPostsToDisplay();
    });
  }

  setPostsToDisplay = () => {
    this.postsToShow = [];
    if (this.authService.isLoggedIn && this.showPublic) {
      this.postsToShow = this.joinArraysWithNoCopies(this.posts, this.publicPosts);
    }
    if (this.authService.isLoggedIn && !this.showPublic) {
      this.postsToShow = this.posts;
    }
    else {
      this.postsToShow = this.publicPosts;
    }
    this.doFilter();
  }

  doFilter = () => {
    if (this.tagForSearch.trim().length > 0) {
      this.postsToShow = this.postsToShow.filter(p => (p.hashtag === this.tagForSearch));
    }
    if (this.locationForSearch.trim().length > 0) {
      this.postsToShow = this.postsToShow.filter(p => (p.location === this.locationForSearch));
    }
    let id = localStorage.getItem('id');
    if (id) {
      this.postsToShow = this.postsToShow.filter(p => {
        let flag = true;
        for (let rep of p.reports) {
          if (rep.userid === id) {
            flag = false;
            break;
          }
        }
        return flag;
      });
    }
    this.fetchCommentData();
    this.fetchLikeData();
    this.fetchDislikeData();
  }

  fetchCommentData = () => {
    if (this.postsToShow) {
      for (let p of this.postsToShow) {
        for (let c of p.comments) {
          if (!this.userData.get(c.userid)) {
            this.userService.getUsersById(c.userid).subscribe(res => {
              this.userData.set(c.userid, res);
            });
          }
        }
      }
    }
  }

  fetchLikeData = () => {
    if (this.postsToShow) {
      for (let p of this.postsToShow) {
        for (let c of p.likes) {
          if (!this.userData.get(c.userid)) {
            this.userService.getUsersById(c.userid).subscribe(res => {
              this.userData.set(c.userid, res);
            });
          }
        }
      }
    }
  }

  fetchDislikeData = () => {
    if (this.postsToShow) {
      for (let p of this.postsToShow) {
        for (let c of p.dislikes) {
          if (!this.userData.get(c.userid)) {
            this.userService.getUsersById(c.userid).subscribe(res => {
              this.userData.set(c.userid, res);
            });
          }
        }
      }
    }
  }

  joinArraysWithNoCopies = (array1: any[], array2: any[]): any[] => {
    let retVal: any[] = [];
    retVal = array1;
    for (let i = 0; i < array2.length; i++) {
      let flag = true;
      for (let j = 0; j < array1.length; j++) {
        if (array2[i].id === array1[j].id) {
          flag = false;
          break;
        }
      }
      if (flag) {
        retVal.push(array2[i]);
      }
    }
    return retVal;
  }

  getCurrentUserImage = (): string => {
    if(!this.authService.isLoggedIn) return "";
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
      this.postsService.leaveComment({ ownerid: post.userid, postid: post.id, userid: id, comment: commentValue }).subscribe(async res => {
        let flag = await this.isNotificationNeededComments(this.userData.get(post.userid), this.currentUser);
          if (flag) {
            this.notifyService.notifyUser(
              {
                "userId": id,
                "notify_user_id": post.userid,
                "type_of_notify": 'comment',
                "notify_id": post.id
              }
            ).subscribe(res => res);
          }
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

  isNotificationNeededLikes = async (user: any, myAccount: any): Promise<boolean> => {
    if (!user.notifyLike) {
      return false;
    }
    if (user.notifyLike && user.notifyLikeFromNotFollowProfile) {
      return true;
    }
    let result : any = await this.followService.checkFollowing(myAccount.id, user.id);  
    return result.isFollowing;

  }

  isNotificationNeededDislikes = async (user: any, myAccount: any): Promise<boolean> => {
    if (!user.notifyDislike) {
      return false;
    }
    if (user.notifyDislike && user.notifyDislikeFromNotFollowProfile) {
      return true;
    }
    let result : any = await this.followService.checkFollowing(myAccount.id, user.id);  
    return result.isFollowing;

  }

  isNotificationNeededComments = async (user: any, myAccount: any): Promise<boolean> => {
    if (!user.notifyComments) {
      return false;
    }
    if (user.notifyComments && user.notifyCommentFromNotFollowProfile) {
      return true;
    }
    let result : any = await this.followService.checkFollowing(myAccount.id, user.id);  
    return result.isFollowing;

  }

}
