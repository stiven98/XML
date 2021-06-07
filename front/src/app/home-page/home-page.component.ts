import { Component, OnInit } from '@angular/core';
import { PostsService } from '../service/posts.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  constructor(public postsService: PostsService, public userService: UserService) { }

  public posts: any[] = [];

  public userData: Map<string, any> = new Map<string, any>();

  convertTimeToText = (timeString : string) => {
    let time = new Date(timeString);
    let now = new Date();
    let diff = now.valueOf() - time.valueOf();
    if(diff >= 86400000){
      return Math.ceil(diff/86400000) + ' days ago';
    }
    if(diff >= 3600000){
      return Math.ceil(diff/3600000) + ' hours ago';
    }
    if(diff >= 60000){
      return Math.ceil(diff/60000) + ' minutes ago';
    }
    else{
      return Math.ceil(diff/1000) + ' seconds ago';
    }
  }

  likeClick = (post : any) => {
    let postId = post.id;
    let userId = post.userid;
  }

  ngOnInit(): void {
    let id = localStorage.getItem('id');
    if (id) {
      this.postsService.getFeed(id).subscribe(res => {
        this.posts = res as any[];
        console.log(res);
        for (let p of this.posts) {
          if (!this.userData.get(p.userid)) {
            this.userService.getUsersById(p.userid).subscribe(res => {
                this.userData.set(p.userid, res);
                console.log(res);
            });
          }
        }
      });
    }
  }

}
