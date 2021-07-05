import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { Post } from '../model/Post.model';
import { ConfigService } from './authorization/config.service';

@Injectable({
  providedIn: 'root'
})
export class StoryService {
  getPagedFeed(id: string, pageNum: number, neededRes: number) {
    return this.http.get(this.config.get_story_page_feed, {params : {id : id, pageNumber : pageNum.toString(), neededResults : neededRes.toString()}}).pipe(
      map((item) => {
        return item;
      })
    );
  }

  constructor(private http: HttpClient, private config:ConfigService) { }

  uploadStory = (formData: FormData) => {
    return this.http.post(this.config.upload_story_picture, formData).pipe(
      map((item) => {
        return item;
      })
    );
  };

  createStory = (post: Post) => {
    return this.http
      .post(this.config.create_story, post)
      .pipe((res) => res);
  };

  getFeed = (id: string) => {
    return this.http.get(this.config.get_story_feed + id).pipe(
      map((item) => {
        return item;
      })
    );
  };
}
