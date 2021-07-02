import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { RegularUser } from '../model/RegularUserModel';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';
import { FollowService } from '../service/follow.service';
import { ManagementService } from '../service/management.service';
import { PostsService } from '../service/posts.service';

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

  public posts: any[] = [];

  constructor(
    private route: ActivatedRoute,
    private userService: UserService,
    public authService: AuthService,
    private router: Router,
    private followService: FollowService,
    private managementService: ManagementService,
    private postsService: PostsService,
    private cdr: ChangeDetectorRef
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
      });

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
    this.managementService.isSubscribed(this.myId,this.id).subscribe((res:any) => this.isSubscriber = res)
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

  block = () => {
    this.managementService
      .blockUser(this.myId, this.id)
      .subscribe((res: any) => {this.initData()});
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
    this.managementService.subscribe(this.myId, this.id).subscribe((res:any) => this.isSubscriber = !this.isSubscriber )
  }

  unSubscribe = () => {
    this.managementService.unSubscribe(this.myId, this.id).subscribe((res:any) => this.isSubscriber = !this.isSubscriber )
  }

}
