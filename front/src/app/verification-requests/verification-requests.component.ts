import { Component, OnInit } from '@angular/core';
import {VerificationRequestService} from '../service/verificationRequest.service';
import {UserService} from '../service/user.service';

@Component({
  selector: 'app-verification-requests',
  templateUrl: './verification-requests.component.html',
  styleUrls: ['./verification-requests.component.css']
})
export class VerificationRequestsComponent implements OnInit {

  // tslint:disable-next-line:ban-types
  verificationRequests: any [] | undefined;
  constructor(private verificationRequestService: VerificationRequestService,
              private userService: UserService) { }


  ngOnInit(): void {

    this.verificationRequestService.getAllVerificationRequests().subscribe(response => {
      this.verificationRequests = response;
      console.log(this.verificationRequests);

      // @ts-ignore
      // tslint:disable-next-line:foreign forin
      for (const i in this.verificationRequests) {
        console.log(i);
        // @ts-ignore
        this.userService.getUserById(this.verificationRequests[i].user_id).subscribe((r) => {
          console.log(r);
          // @ts-ignore
          // tslint:disable-next-line:max-line-length
          this.verificationRequests[i] = {...this.verificationRequests[i], firstName: r.system_user.firstName, lastName: r.system_user.lastName};
          // @ts-ignore
          console.log(this.verificationRequests[i]);
        });


    }});

  }

  verify = (item: any) => {
    console.log(item);
    this.userService.verify(item.user_id).subscribe(() => {
      this.verificationRequestService.acceptVerification(item.id).subscribe((response) => {
        this.ngOnInit();
      });
    });
  }

  decline = (item: any) => {
    console.log(item);
    this.verificationRequestService.declineVerification(item.id).subscribe((response) => {
      this.ngOnInit();
    });
  }



}
