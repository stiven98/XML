export class Post {
  UserId: string;
  Items: PostItem[];
  Type: string;
  Description: string;
  Location: string;
  Hashtag: string;
  constructor() {
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

export interface FavPost {
  userid: string;
  postid: string;
  ownerid: string;
  collection : PostCollection
}

export interface PostCollection {
  name : string;
}

export interface CommentReq {
  userid: string;
  postid: string;
  comment: string;
  ownerid: string;
}
