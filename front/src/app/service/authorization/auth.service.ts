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
    private router: Router,
    private http: HttpClient
  ) {}

  private access_token = null;

  signIn(user: AccountInfoModel) {
    return this.apiService.post('https://localhost:8080/auth', user).subscribe(
      (res: any) => {
        localStorage.setItem('access_token', res.token);
        let authority = res.role;
        localStorage.setItem('role', authority);
        localStorage.setItem('id', res.id);
        localStorage.setItem('username', res.username);
        if (authority === 'ROLE_ADMIN') {
          this.router.navigate(['/homePage']);
        } else if (authority === 'ROLE_SYSTEM_USER') {
          this.router.navigate(['/homePage']);
        } else if (authority === 'ROLE_AGENT') {
          this.router.navigate(['/homePage']);
        }
      },
      (error) => {
        if (error.status === 401) alert('Pogresno ime ili lozinka');
      }
    );
  }

  getToken() {
    return localStorage.getItem('access_token');
  }

  sgetRole(): string {
    return localStorage.getItem('role') as string;
  }
  getUsername = () => {
    return localStorage.getItem('username');
  };

  get isLoggedIn(): boolean {
    let authToken = localStorage.getItem('access_token');
    return authToken !== null ? true : false;
  }

  doLogout() {
    let removeToken = localStorage.removeItem('access_token');
    localStorage.removeItem('role');
    localStorage.removeItem('username');
    localStorage.removeItem('id');
    if (removeToken == null) {
      this.router.navigate(['/login']);
    }
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
