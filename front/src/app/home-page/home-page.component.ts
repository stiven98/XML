import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from '../service/authorization/auth.service';
import { StoryService } from '../service/story.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css'],
})
export class HomePageComponent implements OnInit {

  public tagFilter = false;
  public tagForSearch : string = '';
  public locationFilter = false;
  public locationForSearch : string = '';
  public isFirst = true;
  public stories: any[] = [];
  public userData: Map<string, any> = new Map<string, any>();

  constructor(
    public authService: AuthService,
    private router: Router,
    private route: ActivatedRoute,
    private storyService : StoryService,
    private userService : UserService
  ) {}

  ngOnInit(): void {
    let url: string = this.router.url;
    if (url.includes('tag')) {
      this.tagFilter = true;
      this.tagForSearch = this.route.snapshot.params[`tag`];
    }
    if (url.includes('location')) {
      this.locationFilter = true;
      this.locationForSearch = this.route.snapshot.params[`location`];
    }
    this.initData();
  }

  initData = () => {
    let id = localStorage.getItem('id');
      if (id) {
        this.storyService.getFeed(id).subscribe(res => {
          this.stories = res as any[];
          if (this.stories) {
            for (let s of this.stories) {
              if (!this.userData.get(s.userid)) {
                this.userService.getUsersById(s.userid).subscribe(res => {
                  this.userData.set(s.userid, res);
                });
              }
            }
          }
        });
      }
  }

  getClass = () : string => {
    if(this.isFirst){
      this.isFirst = false;
      return 'carousel-item active';
    }
    return 'carousel-item';
  }

  isImage = (name: string): boolean => {
    let imgFormats = ['jpg', 'jpeg', 'gif', 'png', 'tiff', 'bmp'];
    let flag = false;
    for (let i = 0; i < imgFormats.length; i++) {
      if (name.includes(imgFormats[i])) {
        flag = true;
        break;
      }
    }
    return flag;
  }
}
