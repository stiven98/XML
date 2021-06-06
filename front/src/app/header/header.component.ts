import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css'],
})

export class HeaderComponent implements OnInit {
  searchParams: any;
  userId: any;
  usernames: string[] = [];
  usernamesForSearch: string[] = [];
  usernamesToShow: string[] = [];
  isSearchResultVisible: boolean = false;
  constructor(
    public authService: AuthService,
    public usersService: UserService,
    private router:Router
  ) {}

  ngOnInit(): void {
    this.usersService.getAllUsernames().subscribe((response) => {
      this.usernames = response as string[];
      console.log(this.usernames);
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
      this.router.navigate(['/profile/',this.userId])
    });
  }
  onKeyDown = (event: any) => {
    this.isSearchResultVisible = true;
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
}
