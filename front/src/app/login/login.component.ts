import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
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
  email: string = '';
  showForgotPassword: boolean = false;
  constructor(
    private loginService: AuthService,
    private router: Router,
    private userService: UserService,
    private authService: AuthService
  ) {}

  ngOnInit(): void {}

  onLogin = () => {
    if (this.validateInput()) {
      this.loginService.signIn(this.user);
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
    return username.length > 20 ? false : true;
  };
  isValidPassword = (password: string) => {
    return password.length > 20 ? false : true;
  };

  resetInputs = () => {
    this.user.username = '';
    this.user.password = '';
  };

  forgotPassword = () => {
    this.showForgotPassword = true;
  };
  sendMail = () => {
    if (this.validateEmail(this.email)) {
      this.userService.forgotPassword(this.email).subscribe(res => res);
      alert("Proverite Vaš mail!");
      
    } else {
      alert("Neispravno ste uneli email");
      this.email = "";
    }
  };

  validateEmail = (Email: string) => {
    if (Email.match(new RegExp('.+(@).+(.com)'))) {
      return true;
    } else {
      return false;
    }
  };
}
