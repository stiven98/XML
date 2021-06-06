import { Component, OnInit } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Post, PostItem } from '../model/Post.model';
import { PostsService } from '../service/posts.service';
@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  public files: any[] = [];
  public hashtag: string = '';
  public location: string = '';
  public description: string = '';
  constructor(private modalService: NgbModal, private postsService: PostsService) {
  }

  ngOnInit(): void {
  }

  uploadPost(): void {
    let formData = new FormData();
    for (let i = 0; i < this.files.length; i++) {
      formData.append("files", this.files[i]);
    }
    let post : Post = new Post();
    post.Description = this.description;
    post.Location = this.location;
    post.Hashtag = this.hashtag;
    this.postsService.uploadPosts(formData).subscribe(items => {
      console.log(items);
      // TO-DO
      post.Username = "acaaca123";
      post.Items = items as PostItem[];
      this.files.length == 1 ? post.Type = 'post' : post.Type = 'album';
      this.postsService.createPost(post).subscribe(item => item);
      this.modalService.dismissAll();
    });
  }

  onChange(event: any): void {
    for (var i = 0; i < event.target.files.length; i++) {
      this.files.push(event.target.files[i]);
    }
  }

  openModal(modalDial: any) {
    this.files = [];
    this.hashtag = '';
    this.location = '';
    this.description = '';
    this.modalService.open(modalDial, { ariaLabelledBy: 'modal-basic-title' }).result.then((result: any) => {
      console.log(`Closed`);
    }, (reason: any) => {
      console.log(`Dismissed`);
    });
  }

}
