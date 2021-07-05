import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FollowService } from '../service/follow.service';
import { NotifyService } from '../service/notify.service';
import { UserService } from '../service/user.service';

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

  constructor(private followService: FollowService, private userService: UserService, private notifyService: NotifyService, private router: Router) { }

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
    this.notifyService.getAllNotifyByUserID(this.myId).subscribe((res: any) => {
      let list = res;
      this.notify = []
      for (let u of list) {
        this.userService.getUserById(u.userId).subscribe((res: any) => {
          var text = ""
          let postFlag = true;
          if (u.type_of_notify === "like") { text = "lajkovao fotografiju"; }
          if (u.type_of_notify === "comment") { text = "komentarisao fotografiju" }
          if (u.type_of_notify === "dislike") { text = "dislajkovao fotografiju" }
          if (u.type_of_notify === "post") { text = "postavio fotografiju"; u.notify_user_id = u.userId; }
          if (u.type_of_notify === "story") { text = "objavio story"; postFlag = false; }
          if (u.type_of_notify === "message") { text = "poslao poruku"; postFlag = false; }
          let user = {
            "url": res.system_user.picturePath,
            "name": res.system_user.firstName,
            "notifyid": u.notify_id,
            "text": text,
            "userid": u.notify_user_id,
            "isPost": postFlag
          }
          // @ts-ignore
          this.notify?.push(user);
        });
      }
    })
  }


  notificationClick = (userid: string, postid: string) => {
    this.router.navigate(['single-post/' + userid + '/' + postid]);
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
