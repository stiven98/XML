import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { LikeReq, Post } from '../model/Post.model';
import { ReportReq } from '../model/ReportReq';
import { DeletePost } from '../model/DeletePost';

@Injectable({
  providedIn: 'root',
})
export class PostsService {
  constructor(private http: HttpClient) {}

  uploadPosts = (formData: FormData) => {
    return this.http.post('http://localhost:8086/upload', formData).pipe(
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

  getFeed = (id: string) => {
    return this.http.get('http://localhost:8086/posts/feed/' + id).pipe(
      map((item) => {
        return item;
      })
    );
  };

  getPublicPosts = () => {
    return this.http.get('http://localhost:8086/posts/public').pipe(
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

  likePost = (likeReq: LikeReq) => {
    return this.http
      .post('http://localhost:8086/like-post', likeReq)
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
    return this.http.post('http://localhost:8086/posts/delete', deletePost).pipe(
      map((item) => {
        return item;
      })
    );
  }
}
