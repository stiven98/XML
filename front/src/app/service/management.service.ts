import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ManagementService {

  constructor(private http: HttpClient) {
   }

   isBlocked = (blockedById: any, blockedId: any) => {
    return this.http.get('http://localhost:8087/users/isBlocked/' + blockedById + "/" + blockedId).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  isMuted = (mutedById: any, mutedId: any) => {
    return this.http.get('http://localhost:8087/users/isMuted/' + mutedById + "/" + mutedId).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  blockUser  = (blockedById: any, blockedId: any) =>{
    return this.http.post('http://localhost:8087/users/block/' + blockedById + "/" + blockedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  muteUser  = (mutedById: any, mutedId: any) =>{
    return this.http.post('http://localhost:8087/users/mute/' + mutedById + "/" + mutedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  unMuteUser  = (mutedById: any, mutedId: any) =>{
    return this.http.post('http://localhost:8087/users/unmute/' + mutedById + "/" + mutedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }


  unBlockUser  = (blockedById: any, blockedId: any) =>{
    return this.http.post('http://localhost:8087/users/unblock/' + blockedById + "/" + blockedId,"").pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

}
