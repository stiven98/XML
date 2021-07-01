import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { RegularUser } from '../model/RegularUserModel';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';
import {FollowService} from '../service/follow.service';
import { ManagementService } from '../service/management.service';

@Component({
  selector: 'app-profile-page',
  templateUrl: './profile-page.component.html',
  styleUrls: ['./profile-page.component.css']
})
export class ProfilePageComponent implements OnInit {

  id = '';
  status = 'NO_STATUS';
  myId: string | null;

  followers: any | [];
  following: any | [];
  requests: any | undefined;
  user: RegularUser = new RegularUser();
  isMyProfile:boolean = false;
  isBlockedUsesr:boolean = false;
  isMuted:boolean = false;
  showBlockMute: boolean = false;

  constructor(private route: ActivatedRoute, private userService: UserService,
              public authService: AuthService, private router: Router, 
              private followService: FollowService, private managementService:ManagementService) {
    this.id = route.snapshot.params[`id`];
    this.myId = localStorage.getItem('id');
    this.route.paramMap.subscribe(params => {this.ngOnInit(); });
    this.id === this.myId ? this.showBlockMute = false : this.showBlockMute = true;
  }

  ngOnInit(): void {
    this.id === this.myId ? this.showBlockMute = false : this.showBlockMute = true;
    this.initData();
    
  }

  initData = () => {
    this.id === this.myId ? this.isMyProfile = true : this.isMyProfile = false;
    this.id === this.myId ? this.showBlockMute = false : this.showBlockMute = true;
    this.userService.getUserById(this.route.snapshot.params[`id`]).subscribe((response) => {
      this.user = response as RegularUser;
      console.log(this.user);
      console.log(this.isMyProfile)
      if(!this.isMyProfile && !this.user.isPublic) {
        this.isMyProfile = false;
      }
      else if(this.user.isPublic) {
        this.isMyProfile = true;
      }
     
    });

    this.followService.getFollowers(this.id).subscribe(response => {
      this.followers = response;
      console.log('FOLLOWERS:');
      console.log(this.followers);
      if (this.followers.includes(this.myId)) {
        this.status = 'FOLLOW';
      } else {
        this.status = 'NO_FOLLOW';
      }
    });

    this.followService.getFollowing(this.id).subscribe(response => {
      this.following = response;
      console.log(this.following);
    });

    this.followService.getRequests(this.id).subscribe(response => {
      this.requests = response;
      console.log(this.following);
      if (this.requests.includes(this.myId)) {
        this.status = 'REQUEST';
      }
    });

    this.managementService.isBlocked(this.myId,this.id).subscribe((res:any)=> this.isBlockedUsesr = res)
    this.managementService.isMuted(this.myId,this.id).subscribe((res:any)=> this.isMuted = res)
    if(this.followers == undefined) {
      this.followers = []
    }
    if(this.following == undefined) {
      this.following = []
    }
    console.log(this.status)

  }
  onFollow = () => {
    this.followService.follow(this.myId, this.id).subscribe((response) => {
      console.log(response);
      this.initData();
    });
  }

  onUnfollow = () => {
    this.followService.unfollow(this.myId, this.id).subscribe((response) => {
      console.log(response);
      this.initData();
    });
  }


  block = () => {
    this.managementService.blockUser(this.myId,this.id).subscribe((res:any)=> this.isBlockedUsesr = !this.isBlockedUsesr)
  }

  mute = () =>{
    this.managementService.muteUser(this.myId,this.id).subscribe((res:any)=> this.isMuted = !this.isMuted)
  }

  unBlock = () => {
    this.managementService.unBlockUser(this.myId,this.id).subscribe((res:any)=> this.isBlockedUsesr = !this.isBlockedUsesr)
  }

  unMute = () =>{
    this.managementService.unMuteUser(this.myId,this.id).subscribe((res:any)=> this.isMuted = !this.isMuted)
  }


}
