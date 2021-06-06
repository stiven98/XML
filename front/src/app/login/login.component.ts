import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AccountInfoModel } from '../model/AccountInfoModel';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  user: AccountInfoModel = new AccountInfoModel();
  flag: any;
  constructor(
    private loginService: AuthService,
    private router: Router,
    private userService: UserService
  ) {}

  ngOnInit(): void {}

  onLogin = () => {
    if (this.validateInput()) {
      this.userService.isValidLogin(this.user).subscribe((res) => {
        this.flag = res;
        if (this.flag) {
          this.loginService.signIn(this.user);
          this.router.navigate(['/homePage']);
        } else {
          alert('Pogrešno korisničko ime ili lozinka!');
          this.resetInputs();
          return;
        }
      });
    } else {
      this.resetInputs();
      alert('Morate uneti validne podatke za korisničko ime i lozinku');
    }
  };

  validateInput = () => {
    const { username, password } = this.user;
    const input1 = this.isValidUsername(username);
    const input2 = this.isValidPassword(password);
    return input1 && input2;
  };

  isValidUsername = (username: string) => {
    return username.length > 12 ? false : true;
  };
  isValidPassword = (password: string) => {
    return password.length > 12 ? false : true;
  };

  resetInputs = () => {
    this.user.username = '';
    this.user.password = '';
  };
}
