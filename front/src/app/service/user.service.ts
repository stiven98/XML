import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AgetnRegistrationRequest, User } from '../model/User.model';
import { map } from 'rxjs/operators';
import { UserEdit } from '../model/EditUser';
import { BlockUserDTO } from '../model/BlockUser';
import { ConfigService } from './authorization/config.service';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private http: HttpClient, private config:ConfigService) {}

  registrationUser = (user: User) => {
    return this.http.post(this.config.create_users, {
      system_user: {
        ...user,
        DateOfBirth: user.DateOfBirth + 'T01:00:00+01:00',
      },
      PhoneNumber: user.PhoneNumber,
    });
  };
  registerAgent = (agentRegistrationRequest: AgetnRegistrationRequest) => {
    agentRegistrationRequest.DateOfBirth += 'T01:00:00+01:00';
    return this.http.post(this.config.create_agents_request, agentRegistrationRequest);
  };
  createAgent = (agentRegistrationRequest: AgetnRegistrationRequest) => {
    return this.http.post(this.config.create_agents, agentRegistrationRequest);
  };
  createAgentByAdmin = (agentRegistrationRequest: AgetnRegistrationRequest) => {
    agentRegistrationRequest.DateOfBirth += 'T01:00:00+01:00';
    return this.http.post(this.config.create_agents, agentRegistrationRequest);
  };
  declineAgent = (agentRegistrationRequest: AgetnRegistrationRequest) => {
    return this.http.post(this.config.decline_request_agents, agentRegistrationRequest);
  };


  getAllUsernames = () => {
    return this.http.get(this.config.get_all_sys_usersName).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getUserById = (id: string) => {
    return this.http.get(this.config.get_user_by_id + id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };
  getAllAgentRequests = () => {
    return this.http.get(this.config.get_all_request_agents).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  getUserId = (username: string) => {
    return this.http.get(this.config.get_sys_user_by_id + username).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  editUser = (editUser: UserEdit) => {
    return this.http.put(this.config.update_user, editUser).pipe(map((res)=> {return res;}));

  }

  getUsersById = (id: any) => {
    return this.http.get(this.config.get_user_by_id + id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }
  getPublicTags = () => {
    return this.http.get(this.config.get_post_public_tag).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };
  getPublicLocations = () => {
    return this.http.get(this.config.get_post_public_location).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  };

  getSingedInTags = (id:string) => {
    return this.http.get(this.config.get_post_signed_in_tags +  id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }
  getSingedInLocations = (id: string) => {
    return this.http.get(this.config.get_post_signed_in_location + id).pipe(
      map((responseData) => {
        return responseData;
      })
    );
  }

  uploadProfilePicture = (formData: FormData) => {
    return this.http.post(this.config.upload_user_photo, formData).pipe(map(item => {
        return item;
    }));
  }

  verify = (id: string) => {
    return this.http.put(this.config.update_user_verification + id, null)
      .pipe(map((res) => {
        return res;
      }));
  }
  blockUser = (blockUserDto: BlockUserDTO) => {
    return this.http.post(this.config.auth_block_user, blockUserDto).pipe(map(item => {
        return item;
    }));
  }

  getAllUsers = () => {
    return this.http.get('http://profile-service:8085/users/getAll').pipe(map(item => {
      return item;
    }));
  }
}
