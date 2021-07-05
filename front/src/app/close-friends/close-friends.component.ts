import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FollowService } from '../service/follow.service';
import { ManagementService } from '../service/management.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-close-friends',
  templateUrl: './close-friends.component.html',
  styleUrls: ['./close-friends.component.css']
})
export class CloseFriendsComponent implements OnInit {

  public allFollowers : any | [] = [];
  public allFollowersToShow : any[] = [];
  public closeFriends : any[] = [];
  public closeFriendsToShow : any[] = [];
  public id : string = '';
  constructor(private managementService : ManagementService, private userService : UserService, private followerService : FollowService, private router : Router) { }

  ngOnInit(): void {
    this.initData();
  }

  removeClick = (id : any) => {
    this.managementService.removeFromCloseFriends(this.id, id).subscribe(res => {
      this.allFollowers = [];
      this.allFollowersToShow = [];
      this.closeFriends = [];
      this.closeFriendsToShow = [];
      this.initData();
    });
  }

  addClick = (id : any) => {
    this.managementService.addToCloseFriends(this.id, id).subscribe(res => {
      this.allFollowers = [];
      this.allFollowersToShow = [];
      this.closeFriends = [];
      this.closeFriendsToShow = [];
      this.initData();
    });
  }

  imageClick = (id: any) => {
    this.router.navigate(['profile/' + id]);
  };

  initData = () => {
    let id = localStorage.getItem("id");
    if(id){
      this.id = id;
    }
      this.managementService.getCloseFriends(this.id).subscribe(response => {
        this.closeFriends = response as any[];
        if(this.closeFriends){
          for(let friend of this.closeFriends){
            this.userService.getUsersById(friend).subscribe(res => {
              this.closeFriendsToShow.push(res);
            });
          }
        }
        this.followerService.getFollowers(this.id).subscribe((response) => {
          this.allFollowers = response;
          for(let follower of this.allFollowers){
            if(this.closeFriends && this.closeFriends.length > 0){
              if(!this.closeFriends.includes(follower)){
                this.userService.getUsersById(follower).subscribe(res => {
                  this.allFollowersToShow.push(res);
                });
              }
            }else{
              this.userService.getUsersById(follower).subscribe(res => {
                this.allFollowersToShow.push(res);
              });
            }
          }
      });
    });
  }

}
