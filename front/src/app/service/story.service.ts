import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { Post } from '../model/Post.model';
import { ConfigService } from './authorization/config.service';

@Injectable({
  providedIn: 'root'
})
export class StoryService {
  removeHighlight(id: string, storyid: string) {
    return this.http
      .post('http://story-service:8083/story/remove-highlight', {userid: id, storyid: storyid})
      .pipe(map((item) => {
        return item;
      }));
  };

  addToHighlight(id: string, storyid: string) {
    return this.http
      .post('http://story-service:8083/story/highlight', {userid: id, storyid: storyid})
      .pipe(map((item) => {
        return item;
      }));
  };

  getMyPagedHighlights(id: string, pageNum: number, neededRes: number) {
    return this.http.get('http://story-service:8083/story/paged-highlights', {params : {id : id, pageNumber : pageNum.toString(), neededResults : neededRes.toString()}}).pipe(
      map((item) => {
        return item;
      })
    );
  }
  
  getMyPagedStories(id: string, pageNum: number, neededRes: number) {
    return this.http.get('http://story-service:8083/story/my-paged-stories', {params : {id : id, pageNumber : pageNum.toString(), neededResults : neededRes.toString()}}).pipe(
      map((item) => {
        return item;
      })
    );
  }
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
