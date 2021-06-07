import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { RegularUser } from '../model/RegularUserModel';
import { AuthService } from '../service/authorization/auth.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-profile-page',
  templateUrl: './profile-page.component.html',
  styleUrls: ['./profile-page.component.css']
})
export class ProfilePageComponent implements OnInit {
  id: string = '';
  user: RegularUser = new RegularUser();
  constructor(private route: ActivatedRoute, private userService: UserService, public authService: AuthService, private router: Router) {
    this.id = route.snapshot.params[`id`];
    this.route.paramMap.subscribe(params => {this.ngOnInit();})
   }

  ngOnInit(): void {
    this.userService.getUserById(this.route.snapshot.params[`id`]).subscribe((response) => {
      this.user = response as RegularUser;
      console.log(this.user);
    })
  }
}
