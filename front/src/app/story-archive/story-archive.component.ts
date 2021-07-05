import { Component, OnInit } from '@angular/core';
import { Page } from '../model/Post.model';
import { AuthService } from '../service/authorization/auth.service';
import { StoryService } from '../service/story.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-story-archive',
  templateUrl: './story-archive.component.html',
  styleUrls: ['./story-archive.component.css']
})
export class StoryArchiveComponent implements OnInit {

  public stories: any[] = [];
  public userData: Map<string, any> = new Map<string, any>();
  public currentPage = 1;
  public neededResults = 3;
  public totalCount = 0;
  public id = '';
  public storiesHigh: any[] = [];
  public userDataHigh: Map<string, any> = new Map<string, any>();
  public currentPageHigh = 1;
  public neededResultsHigh = 3;
  public totalCountHigh = 0;

  constructor(
    public authService: AuthService,
    private storyService : StoryService,
    private userService : UserService
  ) {}

  ngOnInit(): void {
    this.initData();
    this.initDataHigh();
  }

  initData = () => {
    let id = localStorage.getItem('id') as string;
    this.id = id;
      if (id) {
        this.storyService.getMyPagedStories(id, this.currentPage, this.neededResults).subscribe(res => {
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
  initDataHigh = () => {
    this.storyService.getMyPagedHighlights(this.id, this.currentPageHigh, this.neededResultsHigh).subscribe(res => {
      let page = res as Page;
      this.storiesHigh = page.stories;
      this.totalCountHigh = page.total_count;
      if (this.storiesHigh) {
        for (let st of this.storiesHigh) {
          if (!this.userDataHigh.get(st.userid)) {
            this.userService.getUsersById(st.userid).subscribe(res => {
              this.userDataHigh.set(st.userid, res);
            });
          }
        }
      }
    }, err => {this.currentPageHigh = 1;});
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

  nextPageClickHigh = () => {
    if(this.currentPageHigh < Math.ceil(this.totalCountHigh / this.neededResultsHigh)){
      this.currentPageHigh = this.currentPageHigh + 1;

    } else {
      this.currentPageHigh = 1;
    }
    this.storiesHigh = [];
    this.userDataHigh = new Map<string, any>();
    this.initDataHigh();
  }

  prevPageClickHigh = () => {
    if(this.currentPageHigh > 1){
      this.currentPageHigh = this.currentPageHigh - 1;
    } else{
      this.currentPageHigh = Math.ceil(this.totalCountHigh / this.neededResultsHigh);
    }
    this.storiesHigh = [];
    this.userDataHigh = new Map<string, any>();
    this.initDataHigh();
  }

  addHighlightClick = (story : any) => {
    this.storyService.addToHighlight(this.id, story.id).subscribe(res => {
      this.storiesHigh = [];
      this.userDataHigh = new Map<string, any>();
      this.stories = [];
      this.userData = new Map<string, any>();
      this.currentPage = 1;
      this.currentPageHigh = 1;
      alert("Uspešno ste istakli priču")
      this.initData();
      this.initDataHigh();
    }, 
      err => {alert("Priča je već istaknuta")});
  }

  removeHighlightClick = (story : any) => {
    this.storyService.removeHighlight(this.id, story.id).subscribe(res => {
      this.storiesHigh = [];
      this.userDataHigh = new Map<string, any>();
      this.stories = [];
      this.userData = new Map<string, any>();
      this.currentPage = 1;
      this.currentPageHigh = 1;
      alert("Uspešno ste uklonili priču iz istaknutih priča")
      this.initData();
      this.initDataHigh();
    }, 
      err => {alert("Priča je već uklonjena")});
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
