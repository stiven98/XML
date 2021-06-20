import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {map} from 'rxjs/operators';


@Injectable({
  providedIn: 'root',
})
export class VerificationRequestService {

  constructor(private http: HttpClient) {
  }

  getAllVerificationRequests = () => {
    return this.http.get('http://localhost:8089/verificationRequest/getAll').pipe(map((response) => {
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
    return this.http.post('http://localhost:8089/verificationRequest/' + id, formData).pipe(map((response) => {
      return response;
    }));
  }

  getImage = (image: string) => {
    return this.http.get('http://localhost:8089/images/' + image).pipe(map((response) => {
      return response;
    }));
  }

  acceptVerification = (id: string) => {
    return this.http.put('http://localhost:8089/verificationRequest/accept/' + id, null).pipe(map((response) => {
      return response;
    }));
  }

  declineVerification = (id: string) => {
    return this.http.put('http://localhost:8089/verificationRequest/decline/' + id, null).pipe(map((response) => {
      return response;
    }));
  }
}
