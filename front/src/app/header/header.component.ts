import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Post, PostItem } from '../model/Post.model';
import { PostsService } from '../service/posts.service';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';
import { StoryService } from '../service/story.service';
import { Time } from '@angular/common';
import { Ad, Campaign } from '../model/Campaign';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css'],
})
export class HeaderComponent implements OnInit {
  public files: any[] = [];
  public campaignFiles: any[] = [];

  public hashtag: string = '';
  public location: string = '';
  public description: string = '';
  public website: string = '';
  public males: boolean = false;
  public females: boolean = false;
  public under18: boolean = false;
  public between18and24: boolean = false;
  public between24and35: boolean = false;
  public over35: boolean = false;
  public timesToPlace: number = 1;
  public timeToPlace: string = '';
  public isMultiple: boolean = false;
  public endDate: string = '';
  public startDate: string = '';
  public whenToPlace: string = '';
  public adLink: string = '';
  public currentFile: string = '';
  public ads: Ad[] = [];

  searchParams: any;
  userId: any;
  usernames: string[] = [];
  tags: string[] = [];
  locations: string[] = [];
  usernamesForSearch: string[] = [];
  tagsForSearch: string[] = [];
  tagsToShow: string[] = [];
  locationsForSearch: string[] = [];
  locationsToShow: string[] = [];
  usernamesToShow: string[] = [];
  isSearchResultVisible: boolean = false;
  isUsersSearchSelected: boolean = false;
  isTagsSearchSelected: boolean = false;
  isLocationsSearchSelected: boolean = false;
  isInputDisalbed: boolean = true;
  constructor(
    public authService: AuthService,
    public usersService: UserService,
    private router: Router,
    private modalService: NgbModal,
    private postsService: PostsService,
    private storyService: StoryService
  ) {}

  ngOnInit(): void {
    this.usersService.getAllUsernames().subscribe((response) => {
      this.usernames = response as string[];
      if (this.authService.isLoggedIn) {
        this.usersService
          .getSingedInLocations(localStorage.getItem('id') as string)
          .subscribe((response) => {
            // @ts-ignore
            this.locations = response.keys as string[];
          });
        this.usersService
          .getSingedInTags(localStorage.getItem('id') as string)
          .subscribe((response) => {
            // @ts-ignore
            this.tags = response.keys as string[];
          });
      } else {
        this.usersService.getPublicLocations().subscribe((response) => {
          // @ts-ignore
          this.locations = response.keys as string[];
        });
        this.usersService.getPublicTags().subscribe((response) => {
          // @ts-ignore
          this.tags = response.keys as string[];
        });
      }
    });
  }

  uploadPost(): void {
    let formData = new FormData();
    for (let i = 0; i < this.files.length; i++) {
      console.log(this.files[i]);
      formData.append('files', this.files[i]);
    }
    console.log(formData);
    let post: Post = new Post();
    post.Description = this.description;
    post.Location = this.location;
    post.Hashtag = this.hashtag;
    this.postsService.uploadPosts(formData).subscribe((items) => {
      console.log(items);
      // TO-DO
      let id = localStorage.getItem('id');
      if (id) {
        post.UserId = id;
      }
      post.Items = items as PostItem[];
      this.files.length == 1 ? (post.Type = 'post') : (post.Type = 'album');
      this.postsService.createPost(post).subscribe((item) => item);
      this.modalService.dismissAll();
    });
  }

  uploadCampaign(): void {
    let formData = new FormData();
    for (let i = 0; i < this.ads.length; i++) {
      console.log(this.ads[i]);
      formData.append('files', this.ads[i].path);
    }
    console.log(formData);
    let campaign: Campaign = new Campaign();
    campaign.description = this.description;
    campaign.ismultiple = this.isMultiple;
    campaign.startday = this.startDate + 'T01:00:00+01:00';
    if (campaign.ismultiple) {
      campaign.endday = this.endDate + 'T01:00:00+01:00';
    } else {
      campaign.endday = campaign.startday;
    }
    campaign.showtomen = this.males;
    campaign.showtowomen = this.females;
    campaign.whentoplace = this.whenToPlace;
    campaign.timestoplace = this.timesToPlace;
    campaign.showunder18 = this.under18;
    campaign.show18to24 = this.between18and24;
    campaign.show24to35 = this.between24and35;
    campaign.showover35 = this.over35;

    this.postsService.uploadPosts(formData).subscribe((items) => {
      console.log(items);
      // TO-DO
      let id = localStorage.getItem('id');
      if (id) {
        campaign.userId = id;
      }
      let Items = items as any;
      console.log(Items);
      console.log(this.ads);
      for (let i = 0; i < Items.length; i++) {
        let ad = new Ad();
        ad.link = this.ads[i].link;
        ad.path = Items[i].path;
        console.log(ad);
        campaign.ads.push(ad);
      }
      console.log(campaign);

      campaign.ads.length == 1
        ? (campaign.type = 'post')
        : (campaign.type = 'album');
      this.postsService.createCampaign(campaign).subscribe((item) => item);
      this.resetAllInputs();
      this.modalService.dismissAll();
    });
  }

  resetAllInputs = () => {
    this.description = '';
    this.males = false;
    this.females = false;
    this.under18 = false;
    this.between18and24 = false;
    this.between24and35 = false;
    this.over35 = false;
    this.timesToPlace = 1;
    this.timeToPlace = '';
    this.isMultiple = false;
    this.endDate = '';
    this.startDate = '';
    this.whenToPlace = '';
    this.adLink = '';
    this.currentFile = '';
    this.ads = [];
  };

  uploadStory(): void {
    let formData = new FormData();
    for (let i = 0; i < this.files.length; i++) {
      console.log(this.files[i]);
      formData.append('files', this.files[i]);
    }
    console.log(formData);
    let post: Post = new Post();
    post.Description = this.description;
    post.Location = this.location;
    post.Hashtag = this.hashtag;
    this.storyService.uploadStory(formData).subscribe((items) => {
      console.log(items);
      // TO-DO
      let id = localStorage.getItem('id');
      if (id) {
        post.UserId = id;
      }
      post.Items = items as PostItem[];
      this.files.length == 1 ? (post.Type = 'story') : (post.Type = 'album');
      this.storyService.createStory(post).subscribe((item) => item);
      this.modalService.dismissAll();
    });
  }

  onChange(event: any): void {
    for (var i = 0; i < event.target.files.length; i++) {
      this.files.push(event.target.files[i]);
    }
  }
  onCampaignChange(event: any): void {
    for (var i = 0; i < event.target.files.length; i++) {
      this.campaignFiles.push(event.target.files[i]);
    }
  }
  openModal2(modalDial: any, f: any) {
    console.log(f);
    this.currentFile = f;
    this.modalService
      .open(modalDial, { ariaLabelledBy: 'modal-basic-title' })
      .result.then(
        (result: any) => {
          console.log(`Closed`);
        },
        (reason: any) => {
          console.log(`Dismissed`);
        }
      );
  }

  onAddLinkClick = () => {
    let ad = new Ad();
    ad.link = this.adLink;
    ad.path = this.currentFile;
    this.adLink = '';
    this.ads.push(ad);
    console.log(this.ads);
  };

  openModal(modalDial: any) {
    this.files = [];
    this.hashtag = '';
    this.location = '';
    this.description = '';
    this.modalService
      .open(modalDial, { ariaLabelledBy: 'modal-basic-title' })
      .result.then(
        (result: any) => {
          console.log(`Closed`);
        },
        (reason: any) => {
          console.log(`Dismissed`);
        }
      );
  }

  onLogout = () => {
    this.authService.doLogout();
  };

  onClick = () => {
    this.isSearchResultVisible = false;
  };

  onLinkClick = (username: string) => {
    if (this.isUsersSearchSelected) {
      this.usersService.getUserId(username).subscribe((response) => {
        this.userId = response;
        this.isSearchResultVisible = false;
        this.searchParams = '';
        this.isUsersSearchSelected = false;
        this.isTagsSearchSelected = false;
        this.usernamesToShow = [];
        this.isLocationsSearchSelected = false;
        this.isInputDisalbed = true;
        this.router.navigate(['/profile', this.userId]);
      });
    }
    if (this.isTagsSearchSelected) {
      this.isSearchResultVisible = false;
      this.isUsersSearchSelected = false;
      this.isTagsSearchSelected = false;
      this.usernamesToShow = [];
      this.searchParams = '';
      this.isLocationsSearchSelected = false;
      this.isInputDisalbed = true;
      this.router.navigate(['/homePage/tag/', username]);
    }
    if (this.isLocationsSearchSelected) {
      this.isSearchResultVisible = false;
      this.isUsersSearchSelected = false;
      this.isTagsSearchSelected = false;
      this.usernamesToShow = [];
      this.isLocationsSearchSelected = false;
      this.isInputDisalbed = true;
      this.searchParams = '';
      this.router.navigate(['/homePage/location/', username]);
    }
  };
  myProfileClick = () => {
    this.router.navigate(['/profile', localStorage.getItem('id')]);
  };
  onKeyDown = () => {
    this.isSearchResultVisible = true;
    if (this.searchParams == '') {
      this.isSearchResultVisible = false;
    }
    if (this.isUsersSearchSelected) {
      this.searchUsers();
    }
    if (this.isTagsSearchSelected) {
      this.searchTags();
    }
    if (this.isLocationsSearchSelected) {
      this.searchLocations();
    }
  };
  searchUsers = () => {
    console.log(this.searchParams);
    for (let username of this.usernames) {
      if (username.includes(this.searchParams)) {
        if (!this.usernamesForSearch.includes(username)) {
          this.usernamesForSearch.push(username);
        }
      }
    }
    this.usernamesToShow = this.usernamesForSearch.filter(
      (value, index, arr) => {
        return value.includes(this.searchParams);
      }
    );
  };
  searchTags = () => {
    console.log(this.searchParams);
    for (let tag of this.tags) {
      if (tag.includes(this.searchParams)) {
        if (!this.tagsForSearch.includes(tag)) {
          this.tagsForSearch.push(tag);
        }
      }
    }
    this.usernamesToShow = this.tagsForSearch.filter((value, index, arr) => {
      return value.includes(this.searchParams);
    });
  };

  searchLocations = () => {
    console.log(this.searchParams);
    for (let location of this.locations) {
      if (location.includes(this.searchParams)) {
        if (!this.locationsForSearch.includes(location)) {
          this.locationsForSearch.push(location);
        }
      }
    }
    this.usernamesToShow = this.locationsForSearch.filter(
      (value, index, arr) => {
        return value.includes(this.searchParams);
      }
    );
  };

  setUsersSearchActive = () => {
    this.isUsersSearchSelected = true;
    this.isTagsSearchSelected = false;
    this.isLocationsSearchSelected = false;
    this.isInputDisalbed = false;
  };
  setTagsSearchActive = () => {
    console.log(this.tags);
    this.isUsersSearchSelected = false;
    this.isTagsSearchSelected = true;
    this.isLocationsSearchSelected = false;
    this.usernamesToShow = [];
    this.isInputDisalbed = false;
  };
  setLocationsSearchActive = () => {
    console.log(this.locations);
    this.isUsersSearchSelected = false;
    this.isTagsSearchSelected = false;
    this.isLocationsSearchSelected = true;
    this.usernamesToShow = [];
    this.isInputDisalbed = false;
  };
  onFocus = () => {
    this.isSearchResultVisible = true;
  };
}
