import {Injectable} from '@angular/core';
import { environment } from '../authorization/enviroment';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {

  private _api_url = environment.apiUrl + '/api';
  private _auth_url = environment.apiUrl +'/auth';
  private _refresh_token_url = this._api_url + '/refresh';


  private _user_url = environment.apiUrl_profile_service + '/users';
  private _sys_users_url = environment.apiUrl_profile_service + '/sysusers';

  private _post_url = environment.apiUrl_post_service + '/posts';
  private _auth_url_api = environment.apiUrl_auth_service + '/api';

  private _follower_url = environment.apiUrl_follower_service + '/users'
  private managment_url = environment.apiUrl_managment_service + '/users'

  private admin_url_verifivation_request = environment.apiUrl_admin_service + '/verificationRequest'
  private story_url = environment.apiUrl_story_service + '/story'

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

  private _user_change = this._user_url + '/change';

  get user_change (): string {
    return this._user_change ;
  }
  private _change_password_url = this._user_url + '/changePassword';

  get change_password_url(): string {
    return this._change_password_url;
  }

  //users start
  private _create_users = this._user_url + '/create'
  get create_users (): string {
    return this._create_users;
  }

  private _get_user_by_id = this._user_url + '/getById/'
  get get_user_by_id (): string {
    return this._get_user_by_id;
  }

  private _update_user = this._user_url + '/update'
  get update_user (): string {
    return this._update_user;
  }

  private _update_user_verification = this._user_url + '/updateVerification/'
  get update_user_verification (): string {
    return this._update_user_verification;
  }

  private _upload_user_photo = environment.apiUrl_profile_service + '/upload'
  get upload_user_photo (): string {
    return this._upload_user_photo;
  }
  //users ends


  //sysusers start

  private _get_all_sys_usersName = this._sys_users_url + '/getAllUsernames'
  get get_all_sys_usersName (): string {
    return this._get_all_sys_usersName;
  }

  private _get_sys_user_by_id = this._sys_users_url + '/getUserId/'
  get get_sys_user_by_id (): string {
    return this._get_sys_user_by_id;
  }


  //sysUsers end


  //post start

  private _get_post_public_tag = this._post_url + '/public-tags'
  get get_post_public_tag ():string{
    return this._get_post_public_tag
  }

  private _get_post_public_location = this._post_url + '/public-locations'
  get get_post_public_location ():string{
    return this._get_post_public_location
  }

  private _get_post_signed_in_tags = this._post_url + '/signed-in-tags/'
  get get_post_signed_in_tags ():string{
    return this._get_post_signed_in_tags
  }

  private _get_post_signed_in_location = this._post_url + '/signed-in-locations/'
  get get_post_signed_in_location ():string{
    return this._get_post_signed_in_location
  }


  //post end


  //auth start
  private _auth_block_user = this._auth_url_api + '/blockUser'
  get auth_block_user ():string{
    return this._auth_block_user;
  }

  //auth end




  //follow star
  private _get_followers = this._follower_url + '/getFollowers/'
  get get_followers ():string{
    return this._get_followers
  }

  private _get_following = this._follower_url + '/getFollowing/'
  get get_following ():string{
    return this._get_following
  }

  private _get_following_request = this._follower_url + '/getRequests/'
  get get_following_request ():string{
    return this._get_following_request
  }

  private _follow = this._follower_url + '/follow/'
  get fllow ():string{
    return this._follow
  }

  private _unfollow = this._follower_url + '/unfollow/'
  get unfllow ():string{
    return this._unfollow
  }

  private _approve_follow_request = this._follower_url + '/acceptRequest/'
  get approve_follow_request ():string{
    return this._approve_follow_request
  }
  //follow end


  //managment start

  private _is_blocked = this.managment_url + '/isBlocked/'
  get is_blocked ():string{
    return this._is_blocked;
  }

  private _is_muted = this.managment_url + '/isMuted/'
  get is_muted ():string{
    return this._is_muted;
  }

  private _block = this.managment_url + '/block/'
  get block ():string{
    return this._block;
  }
  
  private _mute = this.managment_url + '/mute/'
  get mute ():string{
    return this._mute;
  }

  private _unmute = this.managment_url + '/unmute/'
  get unmute ():string{
    return this._unmute;
  }

  private _unblock = this.managment_url + '/unblock/'
  get unblock ():string{
    return this._unblock;
  }

  //managment end

  //admin start
    private _get_all_verification_request = this.admin_url_verifivation_request + '/getAll'
    get get_all_verification_request ():string{
        return this._get_all_verification_request;
    }

    private _verify = this.admin_url_verifivation_request + '/'
    get verify ():string{
      return this._verify;
    }
    private _images = environment.apiUrl_admin_service + '/images/';

    get images ():string{
      return this._images
    }
  
    private _accept_verification_request = this.admin_url_verifivation_request + '/accept/'
    get accept_verification_request ():string{
        return this._accept_verification_request;
    }

    private _decline_verification_request = this.admin_url_verifivation_request + '/decline/'
    get decline_verification_request ():string{
        return this._decline_verification_request;
    }


  //admin end

  //story start
    private _upload_story_picture = environment.apiUrl_story_service + '/upload'
    get upload_story_picture():string{
        return this._upload_story_picture
    }

    private _create_story = this.story_url 
    get create_story():string{
        return this._create_story
    }

    
    private _get_story_feed = this.story_url + '/feed/'
    get get_story_feed():string{
        return this._get_story_feed
    }

  //story end

  //post start


  //post end
}

