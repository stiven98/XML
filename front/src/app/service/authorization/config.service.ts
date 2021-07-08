import {Injectable} from '@angular/core';
import { environment } from '../authorization/enviroment';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {

  private _api_url = environment.apiUrl + '/api';
  private _auth_url = environment.apiUrl_auth_service +'/auth';
  private _refresh_token_url = this._api_url + '/refresh';


  private _user_url = environment.apiUrl_profile_service + '/users';
  private _sys_users_url = environment.apiUrl_profile_service + '/sysusers';

  private _post_url = environment.apiUrl_post_service + '/posts';
  private _auth_url_api = environment.apiUrl_auth_service + '/api';

  private _follower_url = environment.apiUrl_follower_service + '/users'
  private managment_url = environment.apiUrl_managment_service + '/users'

  private admin_url_verifivation_request = environment.apiUrl_admin_service + '/verificationRequest'
  private story_url = environment.apiUrl_story_service + '/story'

  private _notify_url = environment.apiUrl_profile_service + '/notify';
  private _campaigns_url = environment.apiUrl_post_service + '/campaigns';

  private _agents_url = environment.apiUrl_profile_service + '/agents';

  get refresh_token_url(): string {
    return this._refresh_token_url;
  }

  private _login_url = this._auth_url ;

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

  private _upload_user_photo = 'http://profile-service:8085' + '/upload'
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

  private _upload_post_img = 'http://post-service:8086' + '/upload'
  get upload_post_img ():string{
    return this._upload_post_img;
  }

  private _create_post = this._post_url + '/create'
  get create_post ():string{
    return this._create_post
  }

  private _get_post_feed = this._post_url + '/feed/'
  get get_post_feed ():string{
    return this._get_post_feed
  }

  private _get_post_public = this._post_url + '/public/'
  get get_post_public ():string{
    return this._get_post_public
  }

  private _get_post_liked = this._post_url + '/liked/'
  get get_post_liked ():string{
    return this._get_post_liked
  }

  private _get_post_disliked = this._post_url + '/disliked/'
  get get_post_disliked ():string{
    return this._get_post_disliked
  }

  private _get_post_reported = this._post_url + '/reported/'
  get get_post_reported ():string{
    return this._get_post_reported
  }

  private _like_post = environment.apiUrl_post_service + '/like-post'
  get like_post ():string{
    return this._like_post
  }

  private _comment_post = environment.apiUrl_post_service + '/comments'
  get comment_post ():string{
    return this._comment_post
  }

    
  private _dislike_post = environment.apiUrl_post_service + '/dislike-post'
  get dislike_post ():string{
    return this._dislike_post
  }


  private _report_post = environment.apiUrl_post_service + '/report-post'
  get report_post ():string{
    return this._report_post
  }

  private _delete_post = this._post_url+ '/delete'
  get delete_post ():string{
    return this._delete_post
  }
  
  private _edit_archived_post = this._post_url+ '/edit-archived'
  get edit_archived_post ():string{
    return this._edit_archived_post
  }
  
  private _get_post_by_user_id = this._post_url+ '/getByUserId/'
  get get_post_by_user_id ():string{
    return this._get_post_by_user_id
  }

  private _get_post_by_id = this._post_url+ '/getById/'
  get get_post_by_id ():string{
    return this._get_post_by_id
  }
  
  private _save_favourite_post = this._post_url+ '/save'
  get save_favourite_post ():string{
    return this._save_favourite_post
  }
  
  private _get_all_archived_post = this._post_url+ '/all-archived/'
  get get_all_archived_post ():string{
    return this._get_all_archived_post
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

  private _check_following = this._follower_url + '/checkFollowing/'
  get check_following ():string{
    return this._check_following
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

  private _close_friend = this.managment_url + '/closeFriend/'
  get close_friend ():string{
    return this._close_friend;
  }


  private _add_close_friend = this.managment_url + '/addCloseFriend/'
  get add_close_friend ():string{
    return this._add_close_friend;
  }

  private _remove_close_friend = this.managment_url + '/removeCloseFriend/'
  get remove_close_friend ():string{
    return this._remove_close_friend;
  }


  private _subsribe = this.managment_url + '/subscribe/'
  get subsribe ():string{
    return this._subsribe;
  }

  private _unsubsribe = this.managment_url + '/unsubscribe/'
  get unsubsribe ():string{
    return this._unsubsribe;
  }

  private _is_subscribed = this.managment_url + '/isSubscribed/'
  get is_subscribed ():string{
    return this._is_subscribed;
  }

  private _get_all_subscriberd = this.managment_url + '/subscribers/'
  get get_all_subscriberd ():string{
    return this._get_all_subscriberd;
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
    private _images = 'http://admin-service:8089' + '/images/';

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
    private _upload_story_picture = 'http://story-service:8083' + '/upload'
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

     
    private _get_story_page_feed = this.story_url + '/paged-feed'
    get get_story_page_feed():string{
        return this._get_story_page_feed
    }

  //story end

  //notify start
  private _get_notify_by_user = this._notify_url + '/getAll/'
  get get_notify_by_user():string{
      return this._get_notify_by_user
  }

  private _create_notify = this._notify_url + '/create'
  get create_notify():string{
      return this._create_notify
  }
  //notify end

  //campaigns start

  private _delete_campaign = this._campaigns_url + '/delete'
  get delete_campaign():string{
      return this._delete_campaign
  }


  private _get_campaign_by_id = this._campaigns_url + '/getById/'
  get get_campaign_by_id():string{
      return this._get_campaign_by_id
  }


  private _update_campaign = this._campaigns_url + '/updateCampaign'
  get update_campaign():string{
      return this._update_campaign
  }

  private _create_campaign = this._campaigns_url + '/createCampaign'
  get create_campaign():string{
      return this._create_campaign
  }

  private _get_user_campaign = this._campaigns_url + '/getUserCampaigns/'
  get get_user_campaign():string{
      return this._get_user_campaign
  }

  //end campaigns


  //start agents
  private _create_agents_request = this._agents_url + '/createRequest'
  get create_agents_request():string{
      return this._create_agents_request
  }

  private _create_agents = this._agents_url + '/create'
  get create_agents():string{
      return this._create_agents
  }

  private _decline_request_agents = this._agents_url + '/declineRequest'
  get decline_request_agents():string{
      return this._decline_request_agents
  }

  private _get_all_request_agents = this._agents_url + '/getAllRequests'
  get get_all_request_agents():string{
      return this._get_all_request_agents
  }

  //end agents



}

