export class Post {
    Username: string;
    Items : PostItem[];
    Type : string;
    Description : string;
    Location : string;
    Hashtag : string;
    constructor() {
        this.Username = '';
        this.Items = [];
        this.Type = 'post'
        this.Description = '';
        this.Location = '';
        this.Hashtag = '';
    }
}

export interface PostItem{
    Id : string;
    Path : string;
}