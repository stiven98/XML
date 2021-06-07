import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { User } from '../model/User.model';
import { map } from 'rxjs/operators';
import { AccountInfoModel } from '../model/AccountInfoModel';
import { RegularUser } from '../model/RegularUserModel';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private http: HttpClient) {}

  registrationUser = (user: User) => {
    return this.http.post('http://localhost:8085/users/create', {
      system_user: {
        ...user,
        DateOfBirth: user.DateOfBirth + 'T01:00:00+01:00',
      },
      PhoneNumber: user.PhoneNumber,
    });
  };

  isValidLogin(user: AccountInfoModel) {
    return this.http
      .get(
        'http://localhost:8080/api/login-details/isValidLogin/' +
          user.username +
          '/' +
          user.password
      )
      .pipe(
        map((res) => {
          return res;
        })
      );
  }

  getAllUsernames = () => {
    return this.http.get('http://localhost:8085/sysusers/getAllUsernames').pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getUserById = (id:string) => {
    return this.http.get('http://localhost:8085/sysusers/getById/'+id ).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getUserId = (username: string) => {
    return this.http.get('http://localhost:8085/sysusers/getUserId/' + username).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };
}
