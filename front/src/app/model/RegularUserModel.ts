class SystemUser {
    Id: string;
    FirstName: string;
    LastName: string;
    Username: string;
    Email: string;
    Password: string;
    Gender: string;
    PhoneNumber: string;
    DateOfBirth: string;
    constructor() {
        this.Id = '';
        this.FirstName = '';
        this.LastName = '';
        this.Username = '';
        this.Email = '';
        this.Password = '';
        this.Gender = '';
        this.PhoneNumber = '';
        this.DateOfBirth = '';
      }
}
export class RegularUser {
    Id: string;
    SystemUser: SystemUser;
    IsPublic: boolean;
    PhoneNumber: string;
    Website: string;
    Biography: string;
    AllowedTags: boolean;
    IsBlocked: boolean;
    
    constructor() {
      this.Id = '';
      this.SystemUser = new SystemUser();
      this.IsPublic = true;
      this.PhoneNumber = '';
      this.Website = '';
      this.Biography = '';
      this.AllowedTags = true;
      this.IsBlocked = false;
    }
  }