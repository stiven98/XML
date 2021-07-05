import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { ConfigService } from './authorization/config.service';
@Injectable({
  providedIn: 'root'
})
export class ManagementService {

  constructor(private http: HttpClient, private config:ConfigService) {
   }


   getCloseFriends(id: string) {
    return this.http.get(this.config.close_friend + id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

   isBlocked = (blockedById: any, blockedId: any) => {
    return this.http.get(this.config.is_blocked + blockedById + "/" + blockedId).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  isMuted = (mutedById: any, mutedId: any) => {
    return this.http.get(this.config.is_muted + mutedById + "/" + mutedId).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  blockUser  = (blockedById: any, blockedId: any) =>{
    return this.http.post(this.config.block + blockedById + "/" + blockedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  muteUser  = (mutedById: any, mutedId: any) =>{
    return this.http.post(this.config.mute + mutedById + "/" + mutedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  addToCloseFriends  = (userId: any, friendId: any) =>{
    return this.http.post(this.config.add_close_friend + userId + "/" + friendId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  removeFromCloseFriends  = (userId: any, friendId: any) =>{
    return this.http.post(this.config.remove_close_friend + userId + "/" + friendId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  unMuteUser  = (mutedById: any, mutedId: any) =>{
    return this.http.post(this.config.unmute + mutedById + "/" + mutedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }


  unBlockUser  = (blockedById: any, blockedId: any) =>{
    return this.http.post(this.config.unblock + blockedById + "/" + blockedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }


  subscribe  = (subscribedById: any, subscribedId: any) =>{
    return this.http.post(this.config.subsribe + subscribedById + "/" + subscribedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  unSubscribe  = (subscribedById: any, subscribedId: any) =>{
    return this.http.post(this.config.unsubsribe + subscribedById + "/" + subscribedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  isSubscribed = (subscribedById: any, subscribedId: any) => {
    return this.http.get(this.config.is_subscribed + subscribedById + "/" + subscribedId).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getAllSubscribed = (id: any) => {
    return this.http.get(this.config.get_all_subscriberd + id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };
}
