import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {map, tap} from 'rxjs/operators';
import {Observable} from 'rxjs';
import { ConfigService } from './authorization/config.service';

@Injectable({
  providedIn: 'root',
})
export class FollowService {
  constructor(private http: HttpClient, private config:ConfigService) {
  }

  getFollowers(id: string): Observable<string []> {
    return this.http
      .get<string[]>(this.config.get_followers + id);
      // .pipe(
        // tap((responseData) => {
        //   console.log('dole');
        //   console.log(responseData);
        //   console.log(responseData.valueOf().hasOwnProperty('keys'));
        //
        //   if (responseData.valueOf().hasOwnProperty(`keys`)) {
        //     console.log(responseData.valueOf());
        //   }
        //   return [];

       // }));
  }

  getFollowing = (id: string) => {
    return this.http
      .get(this.config.get_following + id)
      .pipe(
        map((responseData) => {
          // @ts-ignore
          return responseData.keys;
        }));
  }

  getRequests = (id: string | null) => {
    return this.http
      .get(this.config.get_following_request + id)
      .pipe(
        map((responseData) => {
          // @ts-ignore
          return responseData.keys;
        }));
  }

  follow = (userID: string | null, targetID: string) => {
    return this.http
      .post(this.config.fllow + userID + '/' + targetID, null);
  }

  unfollow = (userID: string | null, targetID: string | null) => {
    return this.http
      .post(this.config.unfllow + userID + '/' + targetID, null);
  }


  approveRequest = (userElement: never, myId: string | null) => {
    return this.http
      .post(this.config.approve_follow_request + userElement + '/' + myId, null);
  }
}
