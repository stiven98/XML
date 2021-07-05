import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { ConfigService } from './authorization/config.service';

@Injectable({
  providedIn: 'root'
})
export class NotifyService {

  constructor(private http: HttpClient, private config: ConfigService) { }



  getAllNotifyByUserID = (id:any) => {
    return this.http.get(this.config.get_notify_by_user + id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }


  notifyUser = (notifyRequest : any) => {
    return this.http.post(this.config.create_notify, notifyRequest).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }


}
