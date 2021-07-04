export class Campaign {
  userId: string;
  influencers: [];
  ads: Ad[];
  type: string;
  description: string;
  ismultiple: boolean;
  startday: string;
  endday: string;
  timestoplace: number;
  whentoplace: string;
  comments: Comment[];
  likes: Like[];
  dislikes: Dislike[];
  timesplaced: number;
  timesclicked: number;
  showtomen: boolean;
  showtowomen: boolean;
  showunder18: boolean;
  show18to24: boolean;
  show24to35: boolean;
  showover35: boolean;
  constructor() {
    this.userId = '';
    this.influencers = [];
    this.ads = [];
    this.type = '';
    this.description = '';
    this.ismultiple = false;
    this.startday = '';
    this.endday = '';
    this.timestoplace = 0;
    this.whentoplace = '';
    this.likes = [];
    this.dislikes = [];
    this.comments = [];
    this.timesplaced = 0;
    this.timesclicked = 0;
    this.showtomen = false;
    this.showtowomen = false;
    this.showunder18 = false;
    this.show18to24 = false;
    this.show24to35 = false;
    this.showover35 = false;
  }
}

export class Ad {
  path: any;
  link: string;
  constructor() {
    this.path = null;
    this.link = '';
  }
}

export interface PostItem {
  Id: string;
  Path: string;
}

export interface Like {
  id: string;
}
export interface Dislike {
  id: string;
}
export interface Comment {
  id: string;
  userid: string;
  timestamp: string;
  value: string;
}
