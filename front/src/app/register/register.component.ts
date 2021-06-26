import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgetnRegistrationRequest, User } from '../model/User.model';
import {UserValidationModel} from '../model/UserValidation.model';
import { AuthService } from '../service/authorization/auth.service';
import { FollowService } from '../service/follow.service';
import {UserService} from '../service/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  user: User = new User();
  agentRegistrationRequest: AgetnRegistrationRequest = new AgetnRegistrationRequest();
  validation: UserValidationModel = new UserValidationModel();
  isAgentRegistration: boolean = false;
  agentReq: AgetnRegistrationRequest = new AgetnRegistrationRequest();
  constructor(private userService: UserService, private followerService: FollowService,
    public authService: AuthService, private router: Router) { }

  ngOnInit(): void {
    if(this.authService.getRole() === 'ROLE_ADMIN') {
      this.isAgentRegistration = true;
    }
  }

  createAccount = () => {
    if(this.authService.getRole() ==='ROLE_ADMIN' && this.isValidForm()) {
      this.agentRegistrationRequest = this.mapUserToRequest(this.user)
      this.userService.createAgentByAdmin(this.agentRegistrationRequest).subscribe(response =>{
        response
        alert("Agent uspesno registrovan!")
        this.router.navigate(['/homePage']);
      });
      return;
    }
    if(this.isAgentRegistration && this.isValidForm()) {
      this.agentRegistrationRequest = this.mapUserToRequest(this.user)
      this.userService.registerAgent(this.agentRegistrationRequest).subscribe(response => {
        alert('Zahtev uspesno poslat');
        this.user = new User();
        this.validation = new UserValidationModel();
        this.agentRegistrationRequest.WebsiteLink = "";
      }, (error => {
        alert('Greska vec postoji agent!');
      }));
        return;
    }
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
  mapUserToRequest = (user: User) => {
    this.agentReq.FirstName = user.FirstName;
    this.agentReq.LastName = user.LastName;
    this.agentReq.Username = user.Username;
    this.agentReq.Email = user.Email;
    this.agentReq.Password = user.Password;
    this.agentReq.Gender = user.Gender;
    this.agentReq.PhoneNumber = user.PhoneNumber;
    this.agentReq.TypeOfUser = "";
    this.agentReq.DateOfBirth = user.DateOfBirth;
    this.agentReq.PicturePath = "";
    this.agentReq.IsApproved = false;
    return this.agentReq;
  }
  registerAsAgent = () => {
    this.isAgentRegistration = !this.isAgentRegistration;
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
