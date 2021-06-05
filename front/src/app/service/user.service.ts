import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {User} from '../model/User.model';

@Injectable({
  providedIn: 'root',
})

export class UserService {
  constructor(private http: HttpClient) {
  }

  registrationUser = (user: User) => {
    return this.http.post('http://localhost:8085/users/create', {
      system_user: {...user, DateOfBirth: user.DateOfBirth + 'T01:00:00+01:00'}, PhoneNumber: user.PhoneNumber
    });
  }
}
