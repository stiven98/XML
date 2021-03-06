import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { CommentReq, FavPost, LikeReq, Post } from '../model/Post.model';
import { ReportReq } from '../model/ReportReq';
import { DeletePost } from '../model/DeletePost';
import { Campaign } from '../model/Campaign';
import { CampaignRequest } from '../model/CampaignRequest';
import { DeleteReq } from '../model/DeleteReq';
import { AddInfluencer } from '../model/AddInfluencer';

@Injectable({
  providedIn: 'root',
})
export class PostsService {
  constructor(private http: HttpClient) {}

  changeCollection = (changeCollectionReq: FavPost) => {
    return this.http
      .post('http://localhost:8086/posts/edit-archived', changeCollectionReq)
      .pipe((res) => res);
  };

  uploadPosts = (formData: FormData) => {
    return this.http.post('http://localhost:8086/upload', formData).pipe(
      map((item) => {
        return item;
      })
    );
  };

  createCampaignRequest = (campaignRequest: CampaignRequest) => {
    return this.http
      .post('http://localhost:8086/campaigns/createRequest', campaignRequest)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  getByUserId = (id: string) => {
    return this.http.get('http://localhost:8086/posts/getByUserId/' + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getCampaignsByUserId = (id: string) => {
    return this.http
      .get('http://localhost:8086/campaigns/getUserCampaigns/' + id)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  getCampaignsByInfluencerId = (id: string) => {
    return this.http
      .get('http://localhost:8086/campaigns/getInfluencerCampaigns/' + id)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  getAngageRequests = (id: string) => {
    return this.http
      .get('http://localhost:8086/campaigns/getUserCampaignReqs/' + id)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };

  createPost = (post: Post) => {
    return this.http
      .post('http://localhost:8086/posts/create', post)
      .pipe((res) => res);
  };
  createCampaign = (campaign: Campaign) => {
    return this.http
      .post('http://localhost:8086/campaigns/createCampaign', campaign)
      .pipe((res) => res);
  };
  updateCampaign = (campaign: Campaign) => {
    return this.http
      .post('http://localhost:8086/campaigns/updateCampaign', campaign)
      .pipe((res) => res);
  };
  addInfluencer = (addInfluencer: AddInfluencer) => {
    return this.http
      .post('http://localhost:8086/campaigns/addInfluencer', addInfluencer)
      .pipe((res) => res);
  };

  getFeed = (id: string) => {
    return this.http.get('http://localhost:8086/posts/feed/' + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getByIds = (userid: string, postid: string) => {
    return this.http
      .get('http://localhost:8086/posts/getById/' + userid + '/' + postid)
      .pipe(
        map((item) => {
          return item;
        })
      );
  };
  getCampaignByIds = (userid: string, campaignid: string) => {
    return this.http
      .get(
        'http://localhost:8086/campaigns/getById/' + userid + '/' + campaignid
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
        'http://localhost:8086/campaigns/getByInfluencerId/' +
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
    return this.http.get('http://localhost:8086/posts/public/' + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getLikedPosts = (id: string) => {
    return this.http.get('http://localhost:8086/posts/liked/' + id).pipe(
      map((item) => {
        return item;
      })
    );
  };
  getDislikedPosts = (id: string) => {
    return this.http.get('http://localhost:8086/posts/disliked/' + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getReportedPosts = () => {
    return this.http.get('http://localhost:8086/posts/reported/').pipe(
      map((item) => {
        return item;
      })
    );
  };

  addFavourite = (favouriteReq: LikeReq) => {
    return this.http
      .post('http://localhost:8086/posts/save', favouriteReq)
      .pipe((res) => res);
  };

  getFavourite = (id: string) => {
    return this.http
      .get('http://localhost:8086/posts/all-archived/' + id)
      .pipe((res) => res);
  };

  likePost = (likeReq: LikeReq) => {
    return this.http
      .post('http://localhost:8086/like-post', likeReq)
      .pipe((res) => res);
  };

  leaveComment = (commentReq: CommentReq) => {
    return this.http
      .post('http://localhost:8086/comments', commentReq)
      .pipe((res) => res);
  };

  dislikePost = (likeReq: LikeReq) => {
    return this.http
      .post('http://localhost:8086/dislike-post', likeReq)
      .pipe((res) => res);
  };
  reportPost = (reportReq: ReportReq) => {
    return this.http
      .post('http://localhost:8086/report-post', reportReq)
      .pipe((res) => res);
  };
  deletePost(deletePost: DeletePost) {
    return this.http
      .post('http://localhost:8086/posts/delete', deletePost)
      .pipe(
        map((item) => {
          return item;
        })
      );
  }

  deleteCampaign(deletePost: DeletePost) {
    return this.http
      .post('http://localhost:8086/campaigns/delete', deletePost)
      .pipe(
        map((item) => {
          return item;
        })
      );
  }
  deleteCampaignReq(deleteReq: DeleteReq) {
    return this.http
      .post('http://localhost:8086/campaigns/deleteReq', deleteReq)
      .pipe(
        map((item) => {
          return item;
        })
      );
  }
}
