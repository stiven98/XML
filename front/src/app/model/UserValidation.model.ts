
export class UserValidationModel {
  validUsername: string;
  validFirstName: string;
  validLastName: string;
  validConfirmPassword: string;
  validPassword: string;
  validEmail: string;
  validPhoneNumber: string;
  validAddress: string;
  validNumber: string;
  validGender: string;
  validDateOfBirthday: string;

  constructor() {
    this.validFirstName = 'no-validate';
    this.validLastName = 'no-validate';
    this.validUsername = 'no-validate';
    this.validEmail = 'no-validate';
    this.validPassword = 'no-validate';
    this.validConfirmPassword = 'no-validate';
    this.validPhoneNumber = 'no-validate';
    this.validAddress = 'no-validate';
    this.validNumber = 'no-validate';
    this.validGender = 'no-validate';
    this.validDateOfBirthday = 'no-validate';
  }

  validateFirstName = (firstName: string) => {
    if (firstName.length > 0) {
      this.validFirstName = 'is-valid';
      return true;
    } else {
      this.validFirstName = 'is-invalid';
      return false;
    }
  }


  validateLastName = (LastName: string) => {
    if (LastName.length > 0) {
      this.validLastName = 'is-valid';
      return true;
    } else {
      this.validLastName = 'is-invalid';
      return false;
    }
  }

  validateUsername = (Username: string) => {
    if (Username.length > 0) {
      this.validUsername = 'is-valid';
      return true;
    } else {
      this.validUsername = 'is-invalid';
      return false;
    }
  }

  validateEmail = (Email: string) => {
    if (Email.match(new RegExp('.+(@).+(.com)'))) {
      this.validEmail = 'is-valid';
      return true;
    } else {
      this.validEmail = 'is-invalid';
      return false;
    }

  }

  validateGender = (Gender: string) => {
    if (Gender !== 'Pol') {
      this.validGender = 'is-valid';
      return true;
    } else {
      this.validGender = 'is-invalid';
      return false;
    }
  }

  validatePassword = (Password: string) => {
    const regex = /(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[@#$%^&+=])(?=\S+$).{6,}/g;
    const isPasswordValidFlag = regex.test(Password);
    if (isPasswordValidFlag) {
      this.validPassword = 'is-valid';
      return true;
    } else {
      this.validPassword = 'is-invalid';
      return false;
    }
  }

  validateConfirmPassword = (Password: string, ConfirmPassword: string) => {
    if (Password === ConfirmPassword) {
      this.validConfirmPassword = 'is-valid';
      return true;
    } else {
      this.validConfirmPassword = 'is-invalid';
      return false;
    }
  }

  validatePhoneNumber = (PhoneNumber: string) => {
    if (PhoneNumber.match(new RegExp('[+][0-9]{3}[-][0-9]{2}[-][0-9]{3}[-][0-9]{2}[-][0-9]{2}')) && PhoneNumber.length === 17) {
      this.validPhoneNumber = 'is-valid';
      return true;
    } else {
      this.validPhoneNumber = 'is-invalid';
      return false;
    }
  }

  validateDateOfBirthday = (DateOfBirthday: string) => {
    if (DateOfBirthday !== 'mm/dd/yyyy') {
      this.validDateOfBirthday = 'is-valid';
      return true;
    } else {
      this.validDateOfBirthday = 'is-invalid';
      return false;
    }
}
}
