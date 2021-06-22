export class Post {
  Id: string;
  UserId: string;
  Items: PostItem[];
  Type: string;
  Description: string;
  Location: string;
  Hashtag: string;
  constructor() {
    this.Id = '';
    this.UserId = '';
    this.Items = [];
    this.Type = 'post';
    this.Description = '';
    this.Location = '';
    this.Hashtag = '';
  }
}

export interface PostItem {
  Id: string;
  Path: string;
}

export interface LikeReq {
  userid: string;
  postid: string;
  ownerid: string;
}

export interface CommentReq {
  userid: string;
  postid: string;
  comment: string;
  ownerid: string;
}
