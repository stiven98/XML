import { HttpClient, HttpHeaders } from '@angular/common/http';
import { map } from 'rxjs/operators'
import { Injectable } from '@angular/core';
import { Post } from '../model/Post.model';

@Injectable({
    providedIn: 'root',
})

export class PostsService {
    constructor(private http: HttpClient) {
    }

    uploadPosts = (formData: FormData) => {
        return this.http.post('http://localhost:8086/upload', formData).pipe(map(item => {
            return item;
        }));
    }

    createPost = (post: Post) => {
        return this.http.post('http://localhost:8086/posts/create', post).pipe(res => res);
    }

    getFeed = (id: string) => {
        return this.http.get('http://localhost:8086/posts/feed/' + id).pipe(map(item => {
            return item;
        }));
    }

}
