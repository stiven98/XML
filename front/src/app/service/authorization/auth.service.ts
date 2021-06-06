import {
  HttpHeaders,
  HttpErrorResponse,
  HttpClient,
} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Observable, throwError } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { AccountInfoModel } from 'src/app/model/AccountInfoModel';
import { UserService } from '../authorization/user.service';
import { ApiService } from './api.service';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  headers = new HttpHeaders().set('Content-Type', 'application/json');
  currentUser = {};
  constructor(
    private apiService: ApiService,
    private config: ConfigService,
    private userService: UserService,
    private router: Router,
    private http: HttpClient
  ) {}

  private access_token = null;

  signIn(user: AccountInfoModel) {
    return this.apiService
      .post('http://localhost:8080/auth', user)
      .subscribe((res: any) => {
        localStorage.setItem('access_token', res.accessToken);
        this.getUserProfile().subscribe((res) => {
          this.currentUser = res;
          let authority = res.role;
          localStorage.setItem('role', authority);
          if (authority === 'ROLE_ADMIN') {
            this.router.navigate(['/homePage']);
          } else if (authority === 'ROLE_SYSTEM_USER') {
            this.router.navigate(['/homePage']);
          } else if (authority === 'ROLE_AGENT') {
            this.router.navigate(['/homePage']);
          }
        });
      });
  }

  getToken() {
    return localStorage.getItem('access_token');
  }

  getRole() {
    return localStorage.getItem('role');
  }

  get isLoggedIn(): boolean {
    let authToken = localStorage.getItem('access_token');
    return authToken !== null ? true : false;
  }

  doLogout() {
    let removeToken = localStorage.removeItem('access_token');
    this.userService.currentUser = {};
    localStorage.removeItem('role');
    if (removeToken == null) {
      this.router.navigate(['/login']);
    }
  }

  // User profile
  getUserProfile(): Observable<any> {
    return this.apiService
      .get(this.config.get_user_url, { headers: this.headers })
      .pipe(
        map((res: Response) => {
          return res || {};
        }),
        catchError(this.handleError)
      );
  }

  tokenIsPresent() {
    return this.access_token != undefined && this.access_token != null;
  }

  handleError(error: HttpErrorResponse) {
    let msg = '';
    if (error.error instanceof ErrorEvent) {
      // client-side error
      msg = error.error.message;
    } else {
      // server-side error
      msg = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    return throwError(msg);
  }
}
