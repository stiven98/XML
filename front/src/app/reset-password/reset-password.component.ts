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
  newPassword: string = '';
  confirmPassword: string = '';
  id: any;
  constructor(
    private route: ActivatedRoute,
    private userService: UserService,
    private router: Router
  ) {
    this.id = route.snapshot.params[`id`];
  }

  ngOnInit(): void {}

  resetPassword = () => {
    console.log(this.newPassword);
    console.log(this.confirmPassword);
    if (this.newPassword === this.confirmPassword) {
      let resetPassword = new ResetPassword();
      resetPassword.password = this.newPassword;
      resetPassword.password2 = this.confirmPassword;
      resetPassword.requestId = this.id;
      this.userService.checkResetPasswordRequest(this.id).subscribe((res) => {
        if (res) {
          alert('Lozinka uspešno promenjena');
          this.userService.resetPassword(resetPassword).subscribe((res) => res);
          this.router.navigate(['/login']);
        } else {
          alert('Došlo je do greške');
          return;
        }
      });
    } else {
      alert('Lozinke se ne poklapaju!');
      this.newPassword = '';
      this.confirmPassword = '';
    }
  };
}
