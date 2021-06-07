import {Injectable} from '@angular/core';
import { environment } from '../authorization/enviroment';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {

  private _api_url = environment.apiUrl + '/api';
  private _auth_url = environment.apiUrl +'/auth';
  private _refresh_token_url = this._api_url + '/refresh';
  private _user_url = this._api_url + '/user';


  get refresh_token_url(): string {
    return this._refresh_token_url;
  }

  private _login_url = this._auth_url + '/login';

  get login_url(): string {
    return this._login_url;
  }

  private _users_url = this._user_url + '/all';

  get users_url(): string {
    return this._users_url;
  }

  private _get_user_url = this._user_url + '/getUser';

  get get_user_url(): string {
    return this._get_user_url;
  }
  private _get_user_by_id = this._user_url + '/getById/';

  get  get_user_by_id (): string {
    return this._get_user_by_id ;
  }
  private _user_change = this._user_url + '/change';

  get user_change (): string {
    return this._user_change ;
  }
  private _change_password_url = this._user_url + '/changePassword';

  get change_password_url(): string {
    return this._change_password_url;
  }
}

