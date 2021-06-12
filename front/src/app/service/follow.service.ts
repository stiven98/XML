import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class FollowService {
  constructor(private http: HttpClient) {
  }

  getFollowers = (id: string) => {
    return this.http
      .get('https://localhost/users/getFollowers/' + id)
      .pipe(
        map((responseData) => {
          // @ts-ignore
          return responseData.keys;
        }));
  }

  getFollowing = (id: string) => {
    return this.http
      .get('https://localhost/users/getFollowing/' + id)
      .pipe(
        map((responseData) => {
          // @ts-ignore
          return responseData.keys;
        }));
  }

  getRequests = (id: string | null) => {
    return this.http
      .get('https://localhost/users/getRequests/' + id)
      .pipe(
        map((responseData) => {
          // @ts-ignore
          return responseData.keys;
        }));
  }

  follow = (userID: string | null, targetID: string) => {
    return this.http
      .post('https://localhost/users/follow/' + userID + '/' + targetID, null);
  }

  unfollow = (userID: string | null, targetID: string | null) => {
    return this.http
      .post('https://localhost/users/unfollow/' + userID + '/' + targetID, null);
  }


  approveRequest = (userElement: never, myId: string | null) => {
    return this.http
      .post('https://localhost/users/acceptRequest/' + userElement + '/' + myId, null);
  }
}
