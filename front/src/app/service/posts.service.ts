import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators'
import { Injectable } from '@angular/core';
import { LikeReq, Post } from '../model/Post.model';

@Injectable({
    providedIn: 'root',
})

export class PostsService {
    constructor(private http: HttpClient) {
    }

    uploadPosts = (formData: FormData) => {
        return this.http.post('https://localhost/upload', formData).pipe(map(item => {
            return item;
        }));
    }

    createPost = (post: Post) => {
        return this.http.post('https://localhost/posts/create', post).pipe(res => res);
    }

    getFeed = (id: string) => {
        return this.http.get('https://localhost/posts/feed/' + id).pipe(map(item => {
            return item;
        }));
    }

    getPublicPosts = () => {
        return this.http.get('https://localhost/posts/public').pipe(map(item => {
            return item;
        }));
    }

    likePost = (likeReq : LikeReq) => {
        return this.http.post('https://localhost/like-post', likeReq).pipe(res => res);
    }

    dislikePost = (likeReq : LikeReq) => {
        return this.http.post('https://localhost/dislike-post', likeReq).pipe(res => res);
    }
}
