import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { RegularUser } from '../model/RegularUserModel';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';
import { FollowService } from '../service/follow.service';
import { ManagementService } from '../service/management.service';
import { PostsService } from '../service/posts.service';
import { Page } from '../model/Post.model';
import { StoryService } from '../service/story.service';

@Component({
  selector: 'app-profile-page',
  templateUrl: './profile-page.component.html',
  styleUrls: ['./profile-page.component.css'],
})
export class ProfilePageComponent implements OnInit {
  id = '';
  status = 'NO_STATUS';
  myId: string | null;

  followers: any | [];
  following: any | [];
  requests: any | undefined;
  user: RegularUser = new RegularUser();
  isMyProfile: boolean = false;
  isBlockedUsesr: boolean = false;
  isMuted: boolean = false;
  showBlockMute: boolean = false;
  isSubscriber: boolean = false;
  public stories: any[] = [];
  public userData: Map<string, any> = new Map<string, any>();
  public currentPage = 1;
  public neededResults = 3;
  public totalCount = 0;
  public posts: any[] = [];
  public campaigns: any[] = [];
  constructor(
    private route: ActivatedRoute,
    private userService: UserService,
    public authService: AuthService,
    private router: Router,
    private followService: FollowService,
    private managementService: ManagementService,
    private postsService: PostsService,
    private cdr: ChangeDetectorRef,
    private storyService : StoryService
  ) {
    this.id = route.snapshot.params[`id`];
    this.myId = localStorage.getItem('id');
    this.route.paramMap.subscribe((params) => {
      this.ngOnInit();
    });
    this.id === this.myId
      ? (this.showBlockMute = false)
      : (this.showBlockMute = true);
    this.route.paramMap.subscribe((params) => {
      this.ngOnInit();
    });
  }

  ngOnInit(): void {
    this.id === this.myId
      ? (this.showBlockMute = false)
      : (this.showBlockMute = true);
    this.initData();
  }

  nextPageClick = () => {
    if(this.currentPage < Math.ceil(this.totalCount / this.neededResults)){
      this.currentPage = this.currentPage + 1;
    } else {
      this.currentPage = 1;
    }
    this.stories = [];
    this.userData = new Map<string, any>();
    this.initData();
  }

  prevPageClick = () => {
    if(this.currentPage > 1){
      this.currentPage = this.currentPage - 1;
    } else{
      this.currentPage = Math.ceil(this.totalCount / this.neededResults);
    }
    this.stories = [];
    this.userData = new Map<string, any>();
    this.initData();
  }

  removeHighlightClick = (story : any) => {
    this.storyService.removeHighlight(this.id, story.id).subscribe(res => {
      this.stories = [];
      this.userData = new Map<string, any>();
      this.currentPage = 1;
      alert("Uspešno ste uklonili priču iz istaknutih priča")
      this.initData();
    }, 
      err => {alert("Priča je već uklonjena")});
  }

  initData = () => {
    this.id === this.myId
      ? (this.isMyProfile = true)
      : (this.isMyProfile = false);
    this.id === this.myId
      ? (this.showBlockMute = false)
      : (this.showBlockMute = true);
    this.userService
      .getUserById(this.route.snapshot.params[`id`])
      .subscribe((response) => {
        this.user = response as RegularUser;
        console.log(this.user);
        console.log(this.isMyProfile);
        if (!this.isMyProfile && !this.user.isPublic) {
          this.isMyProfile = false;
        } else if (this.user.isPublic) {
          this.isMyProfile = true;
        }

        this.postsService.getByUserId(this.id).subscribe((res) => {
          this.posts = res as any[];
        });

        if (this.authService.getRole() === 'ROLE_AGENT') {
            this.postsService.getCampaignsByUserId(this.id).subscribe((result: any) => {
              this.campaigns = result as any[];
            })
        }
      });

      if (this.id) {
        this.storyService.getMyPagedHighlights(this.id, this.currentPage, this.neededResults).subscribe((res : any) => {
          let page = res as Page;
          this.stories = page.stories;
          this.totalCount = page.total_count;
          if (this.stories) {
            for (let s of this.stories) {
              if (!this.userData.get(s.userid)) {
                this.userService.getUsersById(s.userid).subscribe(res => {
                  this.userData.set(s.userid, res);
                });
              }
            }
          }
        }, (err : any) => {this.currentPage = 1;});
      }

    this.followService.getFollowers(this.id).subscribe((response) => {
      this.followers = response;
      console.log('FOLLOWERS:');
      console.log(this.followers);
      if (this.followers.includes(this.myId)) {
        this.status = 'FOLLOW';
      } else {
        this.status = 'NO_FOLLOW';
      }
    });

    this.followService.getFollowing(this.id).subscribe((response) => {
      this.following = response;
      console.log(this.following);
    });

    this.followService.getRequests(this.id).subscribe((response) => {
      this.requests = response;
      console.log(this.following);
      if (this.requests.includes(this.myId)) {
        this.status = 'REQUEST';
      }
    });

    this.managementService
      .isBlocked(this.myId, this.id)
      .subscribe((res: any) => (this.isBlockedUsesr = res));
    this.managementService
      .isMuted(this.myId, this.id)
      .subscribe((res: any) => (this.isMuted = res));
    if (this.followers == undefined) {
      this.followers = [];
    }
    if (this.following == undefined) {
      this.following = [];
    }
    console.log(this.status);
    this.managementService
      .isSubscribed(this.myId, this.id)
      .subscribe((res: any) => (this.isSubscriber = res));
  };
  onFollow = () => {
    this.followService.follow(this.myId, this.id).subscribe((response) => {
      console.log(response);
      this.initData();
    });
  };

  onUnfollow = () => {
    this.followService.unfollow(this.myId, this.id).subscribe((response) => {
      console.log(response);
      this.initData();
    });
  };

  imageClick = (post: any) => {
    this.router.navigate(['single-post/' + post.userid + '/' + post.id]);
  };
  
  openCampaign = (campaign: any) =>{
    this.router.navigate(['single-campaign/' + campaign.userid + '/' + campaign.id])
  }

  block = () => {
    this.managementService
      .blockUser(this.myId, this.id)
      .subscribe((res: any) => {
        this.initData();
      });
  };

  mute = () => {
    this.managementService
      .muteUser(this.myId, this.id)
      .subscribe((res: any) => (this.isMuted = !this.isMuted));
  };

  unBlock = () => {
    this.managementService
      .unBlockUser(this.myId, this.id)
      .subscribe((res: any) => (this.isBlockedUsesr = !this.isBlockedUsesr));
  };

  unMute = () => {
    this.managementService
      .unMuteUser(this.myId, this.id)
      .subscribe((res: any) => (this.isMuted = !this.isMuted));
  };

  subscribe = () => {
    this.managementService
      .subscribe(this.myId, this.id)
      .subscribe((res: any) => (this.isSubscriber = !this.isSubscriber));
  };

  unSubscribe = () => {
    this.managementService
      .unSubscribe(this.myId, this.id)
      .subscribe((res: any) => (this.isSubscriber = !this.isSubscriber));
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
  }
}
