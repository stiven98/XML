import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class NotifyService {

  constructor(private http: HttpClient) { }



  getAllNotifyByUserID = (id:any) => {
    return this.http.get('http://localhost:8085/notify/getAll/' + id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }


  notifyUser = (notifyRequest : any) => {
    return this.http.post('http://localhost:8085/notify/create', notifyRequest).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }


}
