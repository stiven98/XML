class EditUser {
    id: string;
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    password: string;
    gender: string;
    dateOfBirth: string;

    constructor() {
      this.id = '';
      this.firstName = '';
      this.lastName = '';
      this.username = '';
      this.email = '';
      this.password = '';
      this.gender = 'Pol';
      this.dateOfBirth = 'mm/dd/yyyy';
    }
  }

export class UserEdit{
    userId:string;
    system_user:EditUser;
    isPublic:boolean;
    phoneNumber: string;
    webSite:string;
    biography:string;
    allowedTags:boolean;
    isBlocked:boolean
    isVerified
    constructor() {
        this.userId = '';
        this.system_user = new EditUser();
        this.isPublic = false;
        this.phoneNumber = '';
        this.webSite = '';
        this.biography = '';
        this.allowedTags = false;
        this.isBlocked = false;
        this.isVerified = false;
    }
}
