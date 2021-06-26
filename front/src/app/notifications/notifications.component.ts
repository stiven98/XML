import { Component, OnInit } from '@angular/core';
import {FollowService} from '../service/follow.service';
import { NotifyService } from '../service/notify.service';
import {UserService} from '../service/user.service';

@Component({
  selector: 'app-notifications',
  templateUrl: './notifications.component.html',
  styleUrls: ['./notifications.component.css']
})
export class NotificationsComponent implements OnInit {

  requests: any;
  myId: string | null = '';
  users: [] | undefined;
  notify: [] | any

  constructor(private followService: FollowService, private userService: UserService, private notifyService: NotifyService) { }

  ngOnInit(): void {
    this.myId = localStorage.getItem('id');
    this.followService.getRequests(this.myId).subscribe((response) => {
      this.requests = response;
      this.users = [];
      for (const req of this.requests) {
        this.userService.getUserById(req).subscribe((res) => {
          // @ts-ignore
          this.users?.push(res);
        });

      }
    });
    //treba hendlati da li je objavljen stori post
    this.notifyService.getAllNotifyByUserID("12f93d5d-8ef0-40d1-90b9-559285565dd8").subscribe((res:any) => {  
      let list = res;
      this.notify = []
      for (let u of list) {
        console.log(u.userId)
        this.userService.getUserById(u.userId).subscribe((res:any) => {
          console.log(res)
          let user = {
            "url": res.system_user.picturePath,
            "name": res.system_user.firstName,
            "id" : u.notify_id
          }
          // @ts-ignore
          this.notify?.push(user);
          console.log(this.notify)
        });
      }
    })
  }


 

  declineRequest = (event: any, userElement: never) => {
    console.log(userElement);
    this.followService.unfollow(userElement, this.myId).subscribe((response) => {
      console.log(response);
      this.ngOnInit();
    });
  }

  approveRequest = (event: any, userElement: never) => {
    //this.followService.
    this.followService.approveRequest(userElement, this.myId).subscribe((response) => {
      console.log(response);
      this.ngOnInit();
    });
    console.log(userElement);
  }
}
