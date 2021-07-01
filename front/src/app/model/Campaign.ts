export class Campaign {
  id: string;
  items: PostItem[];
  description: string;
  website: string;
  isMultiple: boolean;
  startDate: Date;
  endDate: Date;
  timestoplace: number;
  timetoshow: string;
  targetgroup: string[];
  likes: Like[];
  dislikes: Dislike[];
  comments: Comment[];
  constructor() {
    this.id = '';
    this.items = [];
    this.description = '';
    this.website = '';
    this.isMultiple = false;
    this.startDate = new Date();
    this.endDate = new Date();
    this.timestoplace = 0;
    this.timetoshow = '';
    this.targetgroup = [];
    this.likes = [];
    this.dislikes = [];
    this.comments = [];
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
