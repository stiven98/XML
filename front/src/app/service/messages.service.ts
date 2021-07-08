import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Message} from '../model/Message';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class MessagesService {

  constructor(private http: HttpClient) {}



  sendTextMessage = (id1: string | null, id2: string, message: Message) => {
    return this.http.post('http://localhost/messages/add/' + id1 + '/' + id2, message).pipe(response => {
      return response;
    });
  }

  getConversation = (id1: string | null, id2: string) => {
    return this.http.get('http://localhost/conversations/' + id1 + '/' + id2).pipe(response => {
      return response;
    });
  }
  upload = (formData: FormData) => {
    return this.http.post('http://localhost/images/upload', formData).pipe(
      map((item) => {
        return item;
      })
    );
  };
}
