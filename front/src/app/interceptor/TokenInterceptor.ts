import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpInterceptor
} from '@angular/common/http';
import { AuthService } from '../service/authorization/auth.service';

@Injectable()
export class TokenInterceptor implements HttpInterceptor {
  constructor(public auth: AuthService) {}
  intercept(req: HttpRequest<any>, next: HttpHandler) {
    const authToken = localStorage.getItem('access_token');
    req = req.clone({
        setHeaders: {
            Authorization: "Bearer " + authToken
        }
    });
    return next.handle(req);
  }
}