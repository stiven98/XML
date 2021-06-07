import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { RegularUser } from '../model/RegularUserModel';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-profile-page',
  templateUrl: './profile-page.component.html',
  styleUrls: ['./profile-page.component.css']
})
export class ProfilePageComponent implements OnInit {
  id: string = '';
  user: RegularUser = new RegularUser();
  constructor(private route: ActivatedRoute, private userService: UserService) {
    this.id = route.snapshot.params[`id`];
   }

  ngOnInit(): void {
    this.userService.getUserById(this.id).subscribe((response) => {
      this.user = response as RegularUser;
      console.log(this.user)
    })
  }

}
