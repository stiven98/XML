import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { CommentReq, LikeReq, Post } from '../model/Post.model';
import { ReportReq } from '../model/ReportReq';
import { DeletePost } from '../model/DeletePost';
import { ConfigService } from './authorization/config.service';

@Injectable({
  providedIn: 'root',
})
export class PostsService {
  constructor(private http: HttpClient, private config:ConfigService) {}

  uploadPosts = (formData: FormData) => {
    return this.http.post(this.config.upload_post_img, formData).pipe(
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

  getFeed = (id: string) => {
    return this.http.get(this.config.get_post_feed + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getPublicPosts = () => {
    return this.http.get(this.config.get_post_public).pipe(
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

  likePost = (likeReq: LikeReq) => {
    return this.http
      .post(this.config.like_post, likeReq)
      .pipe((res) => res);
  };

  leaveComment = (commentReq : CommentReq) => {
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
    return this.http.post(this.config.delete_post, deletePost).pipe(
      map((item) => {
        return item;
      })
    );
  }
}
