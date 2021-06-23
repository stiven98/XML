import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {map} from 'rxjs/operators';
import { ConfigService } from './authorization/config.service';


@Injectable({
  providedIn: 'root',
})
export class VerificationRequestService {

  constructor(private http: HttpClient, private config:ConfigService) {
  }

  getAllVerificationRequests = () => {
    return this.http.get(this.config.get_all_verification_request).pipe(map((response) => {
      const ret = [];
      for (const key in response) {
        if (response.hasOwnProperty(key)) {
          // @ts-ignore
          ret.push(response[key]);
        }
      }
      return ret;
    }));
  }

  verify = (id: string | null, formData: FormData) => {
    return this.http.post(this.config.verify + id, formData).pipe(map((response) => {
      return response;
    }));
  }

  getImage = (image: string) => {
    return this.http.get(this.config.images + image).pipe(map((response) => {
      return response;
    }));
  }

  acceptVerification = (id: string) => {
    return this.http.put(this.config.accept_verification_request + id, null).pipe(map((response) => {
      return response;
    }));
  }

  declineVerification = (id: string) => {
    return this.http.put(this.config.decline_verification_request + id, null).pipe(map((response) => {
      return response;
    }));
  }
}
