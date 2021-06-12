import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-activate',
  templateUrl: './activate.component.html',
  styleUrls: ['./activate.component.css']
})
export class ActivateComponent implements OnInit {

  id: any;
  constructor(
    private route: ActivatedRoute,
    private userService: UserService,
    private router: Router
  ) {
    this.id = route.snapshot.params[`id`];
  }

  ngOnInit(): void {
    this.id = this.route.snapshot.params[`id`];
  }

  activateAccount() {
    this.userService.activateAccount(this.id).subscribe((res : any) => {if(res)alert("Uspesno ste aktivirali nalog")});
  }
}
