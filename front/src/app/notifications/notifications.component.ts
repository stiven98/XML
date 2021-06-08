import { Component, OnInit } from '@angular/core';
import {FollowService} from '../service/follow.service';
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

  constructor(private followService: FollowService, private userService: UserService) { }

  ngOnInit(): void {
    this.myId = localStorage.getItem('id');
    this.followService.getRequests(this.myId).subscribe((response) => {
      this.requests = response;
      console.log(response);
      this.users = [];
      for (const req of this.requests) {
        console.log(req);

        this.userService.getUserById(req).subscribe((res) => {
          console.log(res);
          // @ts-ignore
          this.users?.push(res);
          console.log(this.users);
        });

      }
    });

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
