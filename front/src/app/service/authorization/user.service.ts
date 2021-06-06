import { Injectable } from "@angular/core";
import { map } from "rxjs/operators";
import { ApiService } from "./api.service";
import { ConfigService } from "./config.service";

@Injectable({
  providedIn: 'root'
})
export class UserService {

  currentUser = {};
  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }

  initUser() {
    const promise = this.apiService.get(this.config.refresh_token_url).toPromise()
      .then(res => {
        if (res.access_token !== null) {
              this.getMyInfo().toPromise()
            .then(user => {
              this.currentUser = user;
             });
        } 
      })
      .catch(() => null);
    return promise;
  }

  setupUser(user: {}) {
    this.currentUser = user;
  }

  getMyInfo() {
    return this.apiService.get(this.config.get_user_url)
      .pipe(map(user => {
        this.currentUser = user;
        return user;
      }));
  }

  getAll() {
    return this.apiService.get(this.config.users_url);
  }

  getUserByID(id: any){
    return this.apiService.get(this.config.get_user_by_id + id)
    .pipe(map(user => {
      return user;
    }))
  }

  changeAccountInformation(changeUser: any){
    return this.apiService.post(this.config.user_change, changeUser)
    .pipe(map(user => {
      return user;
    }));
  }

  changePassword(passwordsWrapper: any){
    return this.apiService.post(this.config.change_password_url, passwordsWrapper);
  }

}