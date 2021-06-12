import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ResetPassword } from '../model/ResetPassword';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-reset-password',
  templateUrl: './reset-password.component.html',
  styleUrls: ['./reset-password.component.css'],
})
export class ResetPasswordComponent implements OnInit {
  id: any;
  constructor(
    private route: ActivatedRoute,
    private userService: UserService,
    private router: Router
  ) {
    this.id = route.snapshot.params[`id`];
  }

  ngOnInit(): void {}

  activateAccount = () => {
      this.userService.activateAccount(this.id).subscribe((res) => {
          alert('Nalog uspesno aktiviran');
      });
  };
}
