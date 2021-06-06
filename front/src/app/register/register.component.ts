import { Component, OnInit } from '@angular/core';
import { User } from '../model/User.model';
import {UserValidationModel} from '../model/UserValidation.model';
import {UserService} from '../service/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  user: User = new User();
  validation: UserValidationModel = new UserValidationModel();

  constructor(private userService: UserService) { }

  ngOnInit(): void {
  }

  createAccount = () => {
    if (this.isValidForm()) {
      // this.convertDate();
      this.userService.registrationUser(this.user).subscribe(response => {
        alert('KORISNIK REGISTROVAN');
        this.user = new User();
        this.validation = new UserValidationModel();

      }, (error => {
        alert('GRESKA USER VEC POSTOJI!');
      }));
    }

  }

  convertDate = () => {
    const tokens = this.user.DateOfBirth.split('-');
    console.log(tokens);
    // this.user.DateOfBirthday = tokens[]
  }

  isValidForm = ()  => {
      const isValidFirstName =  this.validation.validateFirstName(this.user.FirstName);
      const isValidLastName = this.validation.validateLastName(this.user.LastName);
      const isValidUsername = this.validation.validateUsername(this.user.Username);
      const isValidEmail = this.validation.validateEmail(this.user.Email);

      const isValidGender = this.validation.validateGender(this.user.Gender);
      const isValidPassword = this.validation.validatePassword(this.user.Password);
      const isValidConfirmPassword = this.validation.validateConfirmPassword(this.user.Password, this.user.ConfirmPassword);
      const isValidPhoneNumber = this.validation.validatePhoneNumber(this.user.PhoneNumber);
      const isValidDateOfBirthday = this.validation.validateDateOfBirthday(this.user.DateOfBirth);

      return isValidFirstName && isValidLastName && isValidUsername && isValidEmail &&
        isValidGender && isValidPassword && isValidConfirmPassword && isValidPhoneNumber && isValidDateOfBirthday;
  }

  onKeyDown = () => {
    this.validation = new UserValidationModel();
  }

  changeGender = (event: any) => {
    this.user.Gender = event.target.value;
    console.log(this.user.Gender);
    this.validation = new UserValidationModel();
  }

  changeDate = (event: any) => {
    this.validation = new UserValidationModel();
    console.log(event);
    console.log(event.target.value);
    this.user.DateOfBirth = event.target.value;
  }
}
