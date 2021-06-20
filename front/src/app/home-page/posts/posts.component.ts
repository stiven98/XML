import { Component, Input, OnInit } from '@angular/core';
import { ReportReq } from 'src/app/model/ReportReq';
import { AuthService } from 'src/app/service/authorization/auth.service';
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
  public reportReq: ReportReq = new ReportReq();
  public showPublic: boolean = false;

  constructor(public postsService: PostsService,
    public userService: UserService,
    public authService: AuthService) {
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
    this.postsService.reportPost(this.reportReq).subscribe((res) => res);
  };

  ngOnInit(): void {
    this.initData();
  }

  initData = () => {
    let id = localStorage.getItem('id');
    this.postsService.getPublicPosts().subscribe((res) => {
      this.publicPosts = res as any[];
      for (let p of this.publicPosts) {
        if (!this.userData.get(p.userid)) {
          this.userService.getUsersById(p.userid).subscribe(res => {
            this.userData.set(p.userid, res);
          });
        }
      }
      if (id) {
        this.postsService.getFeed(id).subscribe(res => {
          this.posts = res as any[];
          for (let p of this.posts) {
            if (!this.userData.get(p.userid)) {
              this.userService.getUsersById(p.userid).subscribe(res => {
                this.userData.set(p.userid, res);
              });
            }
          }
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
