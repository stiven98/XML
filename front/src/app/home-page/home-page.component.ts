import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from '../service/authorization/auth.service';

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
  constructor(
    public authService: AuthService,
    private router: Router,
    private route: ActivatedRoute
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
  }
  
}
