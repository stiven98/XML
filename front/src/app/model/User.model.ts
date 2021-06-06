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
