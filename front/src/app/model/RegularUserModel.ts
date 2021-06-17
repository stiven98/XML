class SystemUser {
    id: string;
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    password: string;
    gender: string;
    phoneNumber: string;
    dateOfBirth: string;
    picturePath: string;
    constructor() {
        this.id = '';
        this.firstName = '';
        this.lastName = '';
        this.username = '';
        this.email = '';
        this.password = '';
        this.gender = '';
        this.phoneNumber = '';
        this.dateOfBirth = '';
        this.picturePath = '';
      }
}
export class RegularUser {
    id: string;
    system_user: SystemUser;
    isPublic: boolean;
    phoneNumber: string;
    website: string;
    biography: string;
    allowedTags: boolean;
    isBlocked: boolean;
    
    constructor() {
      this.id = '';
      this.system_user = new SystemUser();
      this.isPublic = true;
      this.phoneNumber = '';
      this.website = '';
      this.biography = '';
      this.allowedTags = true;
      this.isBlocked = false;
    }
  }