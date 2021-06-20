class EditUser {
    id: string;
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    password: string;
    gender: string;
    dateOfBirth: string;
    picturePath: string
  
    constructor() {
      this.id = '';
      this.firstName = '';
      this.lastName = '';
      this.username = '';
      this.email = '';
      this.password = '';
      this.gender = 'Pol';
      this.dateOfBirth = 'mm/dd/yyyy';
      this.picturePath = 'Pol';
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
    isBlocked:boolean;
    acceptMessagesFromNotFollowProfiles:boolean;
    notifyPosts:boolean;
    notifyMessages:boolean;
    notifyStory:boolean;
    notifyComments:boolean;

    constructor() {
        this.userId = '';
        this.system_user = new EditUser();
        this.isPublic = false;
        this.phoneNumber = '';
        this.webSite = '';
        this.biography = '';
        this.allowedTags = false;
        this.isBlocked = false;
        this.acceptMessagesFromNotFollowProfiles = false;
        this.notifyPosts = false;
        this.notifyMessages = false;
        this.notifyStory = false;
        this.notifyComments = false;
    }
}
  