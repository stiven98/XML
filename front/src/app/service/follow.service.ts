import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {map, tap} from 'rxjs/operators';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class FollowService {
  constructor(private http: HttpClient) {
  }

  getFollowers(id: string) {
    return this.http
      .get<string[]>('http://localhost:8088/users/getFollowers/' + id).pipe(
        map((responseData) => {
          // @ts-ignore
          return responseData.keys;
        }));
  }

  getFollowing = (id: string) => {
    return this.http
      .get('http://localhost:8088/users/getFollowing/' + id)
      .pipe(
        map((responseData) => {
          // @ts-ignore
          return responseData.keys;
        }));
  }

  getRequests = (id: string | null) => {
    return this.http
      .get('http://localhost:8088/users/getRequests/' + id)
      .pipe(
        map((responseData) => {
          // @ts-ignore
          return responseData.keys;
        }));
  }

  follow = (userID: string | null, targetID: string) => {
    return this.http
      .post('http://localhost:8088/users/follow/' + userID + '/' + targetID, null);
  }

  unfollow = (userID: string | null, targetID: string | null) => {
    return this.http
      .post('http://localhost:8088/users/unfollow/' + userID + '/' + targetID, null);
  }


  approveRequest = (userElement: never, myId: string | null) => {
    return this.http
      .post('http://localhost:8088/users/acceptRequest/' + userElement + '/' + myId, null);
  }
}
