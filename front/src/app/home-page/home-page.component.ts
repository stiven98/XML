import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ReportReq } from '../model/ReportReq';
import { AuthService } from '../service/authorization/auth.service';
import { PostsService } from '../service/posts.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css'],
})
export class HomePageComponent implements OnInit {
  constructor(
    public postsService: PostsService,
    public userService: UserService,
    public authService: AuthService,
    private router: Router,
    private route: ActivatedRoute
  ) {}

  public posts: any[] = [];
  public publicPosts: any[] = [];
  public postsToShow: any[] = [];
  public userData: Map<string, any> = new Map<string, any>();
  public userDataToShow: Map<string, any> = new Map<string, any>();
  public userDataPublic: Map<string, any> = new Map<string, any>();
  public reportReq: ReportReq = new ReportReq();
  public tagForSearch: string = '';
  public locationForSearch: string = '';

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

  ngOnInit(): void {
    let url: string = this.router.url;
    if (url.includes('tag')) {
      this.tagForSearch = this.route.snapshot.params[`tag`];
    }
    if (url.includes('location')) {
      this.locationForSearch = this.route.snapshot.params[`location`];
    }
    this.initData();
  }

  reportPost = (postId: string, ownerId: string) => {
    this.reportReq.PostId = postId;
    this.reportReq.UserId = localStorage.getItem('id') as string;
    this.reportReq.OwnerId = ownerId;
    this.postsService.reportPost(this.reportReq).subscribe((res) => res);
  };

  initData = () => {
    if (this.tagForSearch == '' && this.locationForSearch == '') {
      let id = localStorage.getItem('id');
      this.postsService.getPublicPosts().subscribe((res) => {
        this.publicPosts = res as any[];
        console.log(res);
        for (let p of this.publicPosts) {
          let flag = true;
          for (let pos of this.postsToShow) {
            if (pos.id == p.id) {
              flag = false;
              break;
            }
          }
          if (flag) this.postsToShow.push(p);
          if (!this.userDataPublic.get(p.userid)) {
            this.userService.getUsersById(p.userid).subscribe((res) => {
              this.userDataPublic.set(p.userid, res);
              this.userDataToShow.set(p.userid, res);
              console.log(res);
            });
          }
        }
        if (id) {
          this.postsService.getFeed(id).subscribe((res) => {
            this.posts = res as any[];
            console.log(res);
            for (let p of this.posts) {
              let flag = true;
              for (let pos of this.postsToShow) {
                if (pos.id == p.id) {
                  flag = false;
                  break;
                }
              }
              if (flag) this.postsToShow.push(p);
              if (!this.userData.get(p.userid)) {
                this.userService.getUsersById(p.userid).subscribe((res) => {
                  this.userData.set(p.userid, res);
                  this.userDataToShow.set(p.userid, res);
                  console.log(res);
                });
              }
            }
          });
        }
      });
    }
    // TAGOVI
    if (this.tagForSearch != '') {
      let id = localStorage.getItem('id');
      this.postsService.getPublicPosts().subscribe((res) => {
        this.publicPosts = res as any[];
        console.log(res);
        for (let p of this.publicPosts) {
          if (p.hashtag.includes(this.tagForSearch)) {
            let flag = true;
            for (let pos of this.postsToShow) {
              if (pos.id == p.id) {
                flag = false;
                break;
              }
            }
            if (flag) this.postsToShow.push(p);
            if (!this.userDataPublic.get(p.userid)) {
              this.userService.getUsersById(p.userid).subscribe((res) => {
                this.userDataPublic.set(p.userid, res);
                this.userDataToShow.set(p.userid, res);
                console.log(res);
              });
            }
          }
        }
        if (id) {
          this.postsService.getFeed(id).subscribe((res) => {
            this.posts = res as any[];
            console.log(res);
            for (let p of this.posts) {
              if (p.hashtag.includes(this.tagForSearch)) {
                let flag = true;
                for (let pos of this.postsToShow) {
                  if (pos.id == p.id) {
                    flag = false;
                    break;
                  }
                }
                if (flag) this.postsToShow.push(p);
                if (!this.userData.get(p.userid)) {
                  this.userService.getUsersById(p.userid).subscribe((res) => {
                    this.userData.set(p.userid, res);
                    this.userDataToShow.set(p.userid, res);
                    console.log(res);
                  });
                }
              }
            }
          });
        }
      });
    }

    // LOKACIJE

    if (this.locationForSearch != '') {
      let id = localStorage.getItem('id');
      this.postsService.getPublicPosts().subscribe((res) => {
        this.publicPosts = res as any[];
        console.log(res);
        for (let p of this.publicPosts) {
          if (p.location.includes(this.locationForSearch)) {
            let flag = true;
            for (let pos of this.postsToShow) {
              if (pos.id == p.id) {
                flag = false;
                break;
              }
            }
            if (flag) this.postsToShow.push(p);
            if (!this.userDataPublic.get(p.userid)) {
              this.userService.getUsersById(p.userid).subscribe((res) => {
                this.userDataPublic.set(p.userid, res);
                this.userDataToShow.set(p.userid, res);
                console.log(res);
              });
            }
          }
        }
        if (id) {
          this.postsService.getFeed(id).subscribe((res) => {
            this.posts = res as any[];
            console.log(res);
            for (let p of this.posts) {
              if (p.location.includes(this.locationForSearch)) {
                let flag = true;
                for (let pos of this.postsToShow) {
                  if (pos.id == p.id) {
                    flag = false;
                    break;
                  }
                }
                if (flag) this.postsToShow.push(p);
                if (!this.userData.get(p.userid)) {
                  this.userService.getUsersById(p.userid).subscribe((res) => {
                    this.userData.set(p.userid, res);
                    this.userDataToShow.set(p.userid, res);
                    console.log(res);
                  });
                }
              }
            }
          });
        }
      });
    }
  };
}
