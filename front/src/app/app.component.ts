import { Component } from '@angular/core';
import { NavigationStart, Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'front';

  showHeader: boolean = false;



  ngOnInit(): void {

  }
  constructor(private router: Router) {
    // on route change to '/login', set the variable showHead to false
      router.events.forEach((event) => {
        if (event instanceof NavigationStart) {
          if (event.url === '/login' || event.url === '/registration' || event.url.includes('reset') ) {
            this.showHeader = false;
          } else {
            // console.log("NU")
            this.showHeader = true;
          }
        }
      });
    }
  }
