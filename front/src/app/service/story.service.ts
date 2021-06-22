import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { Post } from '../model/Post.model';

@Injectable({
  providedIn: 'root'
})
export class StoryService {

  constructor(private http: HttpClient) { }

  uploadStory = (formData: FormData) => {
    return this.http.post('http://localhost:8083/upload', formData).pipe(
      map((item) => {
        return item;
      })
    );
  };

  createStory = (post: Post) => {
    return this.http
      .post('http://localhost:8083/story', post)
      .pipe((res) => res);
  };

  getFeed = (id: string) => {
    return this.http.get('http://localhost:8083/story/feed/' + id).pipe(
      map((item) => {
        return item;
      })
    );
  };
}
