import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { User } from '../model/User.model';
import { map } from 'rxjs/operators';
import { UserEdit } from '../model/EditUser';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private http: HttpClient) {}

  registrationUser = (user: User) => {
    return this.http.post('http://localhost:8085/users/create', {
      system_user: {
        ...user,
        DateOfBirth: user.DateOfBirth + 'T01:00:00+01:00',
      },
      PhoneNumber: user.PhoneNumber,
    });
  };

  getAllUsernames = () => {
    return this.http.get('http://localhost:8085/sysusers/getAllUsernames').pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getUserById = (id:string) => {
    return this.http.get('http://localhost:8085/users/getById/'+id ).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getUserId = (username: string) => {
    return this.http.get('http://localhost:8085/sysusers/getUserId/' + username).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

editUser = (editUser: UserEdit) => {
  return this.http.put('http://localhost:8085/users/update', editUser).pipe(map((res)=> {return res;}));

}

getUsersById = (id: any) => {
  return this.http.get('http://localhost:8085/users/getById/' + id).pipe(
    map((responseData) => {
      return responseData;
    })
  );
}
getPublicTags = () => {
  return this.http.get(' http://localhost:8086/posts/public-tags').pipe(
    map((responseData) => {
      return responseData;
    })
  );
};
getPublicLocations = () => {
  return this.http.get('http://localhost:8086/posts/public-locations').pipe(
    map((responseData) => {
      return responseData;
    })
  );
};

getSingedInTags = (id:string) => {
  return this.http.get('http://localhost:8086/posts/signed-in-tags/' +  id).pipe(
    map((responseData) => {
      return responseData;
    })
  );
};
getSingedInLocations = (id: string) => {
  return this.http.get('http://localhost:8086/posts/signed-in-locations/' + id).pipe(
    map((responseData) => {
      return responseData;
    })
  );
};


}
