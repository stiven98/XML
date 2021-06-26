import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class NotifyService {

  constructor(private http: HttpClient) { }



  getAllNotifyByUserID = (id:string) => {
    return this.http.get('http://localhost:8085/notify/getAll/' + id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }





}
