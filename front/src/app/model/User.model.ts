export class User {
  FirstName: string;
  LastName: string;
  Username: string;
  Email: string;
  Password: string;
  Gender: string;
  PhoneNumber: string;
  DateOfBirth: string;
  ConfirmPassword: string;

  constructor() {
    this.FirstName = '';
    this.LastName = '';
    this.Username = '';
    this.Email = '';
    this.Password = '';
    this.Gender = 'Pol';
    this.PhoneNumber = '';
    this.ConfirmPassword = '';
    this.DateOfBirth = 'mm/dd/yyyy';
  }
}

export class AgetnRegistrationRequest {

  FirstName: string;
  LastName: string;
  Username: string;
  Email: string;
  Password: string;
  Gender: string;
  PhoneNumber: string;
  TypeOfUser: string;
  DateOfBirth: string;
  PicturePath : string;
  WebsiteLink: string;
  IsApproved: boolean;
  constructor() {
    
    this.FirstName = "";
    this.LastName = "";
    this.Username = "";
    this.Email = "";
    this.Password = "";
    this.Gender = "";
    this.PhoneNumber = "";
    this.TypeOfUser = "";
    this.DateOfBirth = "";
    this.PicturePath = "";
    this.WebsiteLink = "";
    this.IsApproved = false;
  }
}
