import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Post, PostItem } from '../model/Post.model';
import { PostsService } from '../service/posts.service';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css'],
})

export class HeaderComponent implements OnInit {
  public files: any[] = [];
  public hashtag: string = '';
  public location: string = '';
  public description: string = '';
  
  searchParams: any;
  userId: any;
  usernames: string[] = [];
  usernamesForSearch: string[] = [];
  usernamesToShow: string[] = [];
  isSearchResultVisible: boolean = false;
  isUsersSearchSelected: boolean = false;
  isTagsSearchSelected: boolean = false;
  isLocationsSearchSelected: boolean = false;
  isInputDisalbed: boolean = true;
  constructor(
    public authService: AuthService,
    public usersService: UserService,
    private router:Router,private modalService: NgbModal, 
    private postsService: PostsService
  ) {}

 

  ngOnInit(): void {
    this.usersService.getAllUsernames().subscribe((response) => {
      this.usernames = response as string[];
      console.log(this.usernames);
    });
  }

  uploadPost(): void {
    let formData = new FormData();
    for (let i = 0; i < this.files.length; i++) {
      formData.append("files", this.files[i]);
    }

    let post : Post = new Post();
    post.Description = this.description;
    post.Location = this.location;
    post.Hashtag = this.hashtag;
    this.postsService.uploadPosts(formData).subscribe(items => {
      console.log(items);
      // TO-DO
      let id = localStorage.getItem('id');
      if(id){
        post.UserId = id;
      }
      post.Items = items as PostItem[];
      this.files.length == 1 ? post.Type = 'post' : post.Type = 'album';
      this.postsService.createPost(post).subscribe(item => item);
      this.modalService.dismissAll();
    });
  }

  onChange(event: any): void {
    for (var i = 0; i < event.target.files.length; i++) {
      this.files.push(event.target.files[i]);
    }
  }

  openModal(modalDial: any) {
    this.files = [];
    this.hashtag = '';
    this.location = '';
    this.description = '';
    this.modalService.open(modalDial, { ariaLabelledBy: 'modal-basic-title' }).result.then((result: any) => {
      console.log(`Closed`);
    }, (reason: any) => {
      console.log(`Dismissed`);
    });
  }

  onLogout = () => {
    this.authService.doLogout();
  };

  onClick = () => {
    this.isSearchResultVisible = false;
  };

  onLinkClick = (username:string) => {
    this.usersService.getUserId(username).subscribe((response) => {
      this.userId = response;
      this.isSearchResultVisible = false;
      this.searchParams = "";
      this.isUsersSearchSelected = false;
      this.isTagsSearchSelected = false;
      this.usernamesToShow = [];
      this.isLocationsSearchSelected = false;
      this.isInputDisalbed = true;
      this.router.navigate(['/profile', this.userId]);
    });
  }
  myProfileClick = () => {
      this.router.navigate(['/profile', localStorage.getItem('id')]);
  }
  onKeyDown = () => {
    this.isSearchResultVisible = true;
    if(this.searchParams == "") {
      this.isSearchResultVisible = false;
    }
    if(this.isUsersSearchSelected) {
      this.searchUsers();
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
  }
  setUsersSearchActive = () => {
    this.isUsersSearchSelected = true;
    this.isTagsSearchSelected = false;
    this.isLocationsSearchSelected = false;
    this.isInputDisalbed = false;
  }
  setTagsSearchActive = () => {
    this.isUsersSearchSelected = false;
    this.isTagsSearchSelected = true;
    this.isLocationsSearchSelected = false;
    this.usernamesToShow = [];
    this.isInputDisalbed = false;
  }
  setLocationsSearchActive = () => {
    this.isUsersSearchSelected = false;
    this.isTagsSearchSelected = false;
    this.isLocationsSearchSelected = true;
    this.usernamesToShow = [];
    this.isInputDisalbed = false;
  }
  onFocus = () => {
    this.isSearchResultVisible = true;
  }
}
