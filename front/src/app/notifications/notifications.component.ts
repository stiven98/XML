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
    this.notifyService.getAllNotifyByUserID(this.myId).subscribe((res:any) => {  
      let list = res;
      this.notify = []
      for (let u of list) {
  
     
        console.log(u.userId)
        this.userService.getUserById(u.userId).subscribe((res:any) => {
          var text = ""
          if(u.type_of_notify === "like"){ text = "lajkova"; console.log("usao  " +  text)}
          if(u.type_of_notify === "comment") {text = "komentarisao"}
          if(u.type_of_notify === "dislike") {text = "dislajkovao"}
          if(u.type_of_notify === "post") {text = "postavio"}
          if(u.type_of_notify === "story") {text = "objavio story"}
          let user = {
            "url": res.system_user.picturePath,
            "name": res.system_user.firstName,
            "id" : u.notify_id,
            "text": text
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
