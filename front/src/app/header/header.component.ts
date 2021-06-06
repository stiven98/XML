import { Component, OnInit } from '@angular/core';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css'],
})
export class HeaderComponent implements OnInit {
  searchParams: any;
  usernames: string[] = [];
  usernamesForSearch: string[] = [];
  usernamesToShow: string[] = [];
  constructor(
    public authService: AuthService,
    public usersService: UserService
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

  onKeyDown = (event: any) => {
    for (let username of this.usernames) {
      if (username.includes(this.searchParams)) {
        if (!this.usernamesForSearch.includes(username)) {
          this.usernamesForSearch.push(username);
        }
      }
    }
  };
}
