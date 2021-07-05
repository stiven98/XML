import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Page } from '../model/Post.model';
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
  public currentPage = 1;
  public neededResults = 3;
  public totalCount = 0;

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
        this.storyService.getPagedFeed(id, this.currentPage, this.neededResults).subscribe(res => {
          let page = res as Page;
          this.stories = page.stories;
          this.totalCount = page.total_count;
          if (this.stories) {
            for (let s of this.stories) {
              if (!this.userData.get(s.userid)) {
                this.userService.getUsersById(s.userid).subscribe(res => {
                  this.userData.set(s.userid, res);
                });
              }
            }
          }
        }, err => {this.currentPage = 1;});
      }
  }

  nextPageClick = () => {
    if(this.currentPage < Math.ceil(this.totalCount / this.neededResults)){
      this.currentPage = this.currentPage + 1;
    } else {
      this.currentPage = 1;
    }
    this.stories = [];
    this.userData = new Map<string, any>();
    this.initData();
  }

  prevPageClick = () => {
    if(this.currentPage > 1){
      this.currentPage = this.currentPage - 1;
    } else{
      this.currentPage = Math.ceil(this.totalCount / this.neededResults);
    }
    this.stories = [];
    this.userData = new Map<string, any>();
    this.initData();
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
