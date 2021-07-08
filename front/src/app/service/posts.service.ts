import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { CommentReq, FavPost, LikeReq, Post } from '../model/Post.model';
import { ReportReq } from '../model/ReportReq';
import { DeletePost } from '../model/DeletePost';
import { ConfigService } from './authorization/config.service';
import { Campaign } from '../model/Campaign';
import { CampaignRequest } from '../model/CampaignRequest';
import { DeleteReq } from '../model/DeleteReq';
import { AddInfluencer } from '../model/AddInfluencer';

@Injectable({
  providedIn: 'root',
})
export class PostsService {
  constructor(private http: HttpClient, private config:ConfigService) {}

  changeCollection = (changeCollectionReq: FavPost) => {
    return this.http
      .post(this.config.edit_archived_post, changeCollectionReq)
      .pipe((res) => res);
  };

  uploadPosts = (formData: FormData) => {
    return this.http.post(this.config.upload_post_img, formData).pipe(
      map((item) => {
        return item;
      })
    );
  };

  createCampaignRequest = (campaignRequest: CampaignRequest) => {
    return this.http
      .post('http://post-service:8086/campaigns/createRequest', campaignRequest)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  getByUserId = (id: string) => {
    return this.http.get(this.config.get_post_by_user_id + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getCampaignsByUserId = (id: string) => {
    return this.http
      .get('http://post-service:8086/campaigns/getUserCampaigns/' + id)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  getCampaignsByInfluencerId = (id: string) => {
    return this.http
      .get('http://post-service:8086/campaigns/getInfluencerCampaigns/' + id)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  getAngageRequests = (id: string) => {
    return this.http
      .get('http://post-service:8086/campaigns/getUserCampaignReqs/' + id)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  createPost = (post: Post) => {
    return this.http
      .post(this.config.create_post, post)
      .pipe((res) => res);
  };
  createCampaign = (campaign: Campaign) => {
    return this.http
      .post(this.config.create_campaign, campaign)
      .pipe((res) => res);
  };
  updateCampaign = (campaign: Campaign) => {
    return this.http
      .post(this.config.update_campaign, campaign)
      .pipe((res) => res);
  };
  addInfluencer = (addInfluencer: AddInfluencer) => {
    return this.http
      .post('http://post-service:8086/campaigns/addInfluencer', addInfluencer)
      .pipe((res) => res);
  };

  getFeed = (id: string) => {
    return this.http.get(this.config.get_post_feed + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getByIds = (userid: string, postid: string) => {
    return this.http
      .get(this.config.get_post_by_id + userid + '/' + postid)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };
  getCampaignByIds = (userid: string, campaignid: string) => {
    return this.http
      .get(
        this.config.get_campaign_by_id + userid + '/' + campaignid
      )
      .pipe(
        map((item) => {
          return item;
        })
      );
  };
  getCampaignByInfluencerIds = (userid: string, campaignid: string) => {
    return this.http
      .get(
        'http://post-service:8086/campaigns/getByInfluencerId/' +
          userid +
          '/' +
          campaignid
      )
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  getPublicPosts = (id = '00000000-00000000-00000000-00000000') => {
    return this.http.get(this.config.get_post_public + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getLikedPosts = (id: string) => {
    return this.http.get(this.config.get_post_liked + id).pipe(
      map((item) => {
        return item;
      })
    );
  };
  getDislikedPosts = (id: string) => {
    return this.http.get(this.config.get_post_disliked + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getReportedPosts = () => {
    return this.http.get(this.config.get_post_reported).pipe(
      map((item) => {
        return item;
      })
    );
  };

  addFavourite = (favouriteReq: LikeReq) => {
    return this.http
      .post(this.config.save_favourite_post, favouriteReq)
      .pipe((res) => res);
  };

  getFavourite = (id: string) => {
    return this.http
      .get(this.config.get_all_archived_post + id)
      .pipe((res) => res);
  };

  likePost = (likeReq: LikeReq) => {
    return this.http
      .post(this.config.like_post, likeReq)
      .pipe((res) => res);
  };

  leaveComment = (commentReq: CommentReq) => {
    return this.http
    .post(this.config.comment_post, commentReq)
    .pipe((res) => res);
};

  dislikePost = (likeReq: LikeReq) => {
    return this.http
      .post(this.config.dislike_post, likeReq)
      .pipe((res) => res);
  };
  reportPost = (reportReq: ReportReq) => {
    return this.http
      .post(this.config.report_post, reportReq)
      .pipe((res) => res);
  };
  deletePost(deletePost: DeletePost) {
    return this.http
      .post(this.config.delete_post, deletePost)
      .pipe(
        map((item) => {
          return item;
        })
      );
  }

  deleteCampaign(deletePost: DeletePost) {
    return this.http
      .post(this.config.delete_campaign, deletePost)
      .pipe(
        map((item) => {
          return item;
        })
      );
  }
  deleteCampaignReq(deleteReq: DeleteReq) {
    return this.http
      .post('http://post-service:8086/campaigns/deleteReq', deleteReq)
      .pipe(
        map((item) => {
          return item;
        })
      );
  }
}
