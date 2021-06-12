import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { User } from '../model/User.model';
import { map } from 'rxjs/operators';
import { UserEdit } from '../model/EditUser';
import { ResetPassword } from '../model/ResetPassword';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private http: HttpClient) {}

  registrationUser = (user: User) => {
    return this.http.post('https://localhost/users/create', {
      system_user: {
        ...user,
        DateOfBirth: user.DateOfBirth + 'T01:00:00+01:00',
      },
      PhoneNumber: user.PhoneNumber,
    });
  };

  getAllUsernames = () => {
    return this.http.get('https://localhost/sysusers/getAllUsernames').pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getUserById = (id:string) => {
    return this.http.get('https://localhost/users/getById/'+id ).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getUserId = (username: string) => {
    return this.http.get('https://localhost/sysusers/getUserId/' + username).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

editUser = (editUser: UserEdit) => {
  return this.http.put('https://localhost/users/update', editUser).pipe(map((res)=> {return res;}));

}

getUsersById = (id: any) => {
  return this.http.get('https://localhost/users/getById/' + id).pipe(
    map((responseData) => {
      return responseData;
    })
  );
}
getPublicTags = () => {
  return this.http.get('https://localhost/posts/public-tags').pipe(
    map((responseData) => {
      return responseData;
    })
  );
};
getPublicLocations = () => {
  return this.http.get('https://localhost/posts/public-locations').pipe(
    map((responseData) => {
      return responseData;
    })
  );
};

getSingedInTags = (id:string) => {
  return this.http.get('https://localhost/posts/signed-in-tags/' +  id).pipe(
    map((responseData) => {
      return responseData;
    })
  );
};
getSingedInLocations = (id: string) => {
  return this.http.get('https://localhost/posts/signed-in-locations/' + id).pipe(
    map((responseData) => {
      return responseData;
    })
  );
};
  forgotPassword = (email: string) => {
    return this.http
      .post('https://localhost/auth/forgotPassword', email)
      .pipe(map((res) => {return res}));
  };

  checkResetPasswordRequest = (id : any) => {
    return this.http
    .get('https://localhost/auth/checkRequest/'+ id)
    .pipe(map((res) => {return res}));
};

  resetPassword = (resetPassword: ResetPassword) => {
    return this.http
    .post('https://localhost/auth/resetPassword', resetPassword)
    .pipe(map((res) => {return res}));
  }

  activateAccount = (id: string) => {
    return this.http
    .post('https://localhost:443/auth/activate', {'id' : id})
    .pipe(map((res) => {return res}));
  }

}
