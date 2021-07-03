import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { FavPost, PostCollection } from '../model/Post.model';
import { PostsService } from '../service/posts.service';

@Component({
  selector: 'app-saved-posts',
  templateUrl: './saved-posts.component.html',
  styleUrls: ['./saved-posts.component.css']
})
export class SavedPostsComponent implements OnInit {

  public allArchivedPosts : any[] = [];
  public collections : any[] = [];
  public posts : any[] = [];
  public postsToShow : any[] = [];
  public collectionInput : string = "";
  public postToChange : any;
  constructor(private postsService : PostsService, private router : Router, private modalService: NgbModal) { }

  ngOnInit(): void {
    this.initData();
  }

  initData = () => {
    let id = localStorage.getItem("id");
    if(id){
      this.postsService.getFavourite(id).subscribe(res => {
        this.allArchivedPosts = res as any[];
        for(let arc of this.allArchivedPosts){
          this.postsService.getByIds(arc.ownerid, arc.postid).subscribe(res => {
            this.posts.push(res as any);
          });
        if(!this.collections.includes(arc.collection.name)){
          this.collections.push(arc.collection.name);
        }
        }
      });
    }
    else{
      this.router.navigate(["/homepage"]);
    }
  }

  doFilter = (collection : string) => {
    this.postsToShow = [];
    for(let post of this.posts){
      for(let favPost of this.allArchivedPosts){
        if(post.id === favPost.postid){
          let col : string = favPost.collection.name;
          if(col.toLowerCase().includes(collection.toLowerCase())){
            this.postsToShow.push(post);
          }
        }
      }
    }
  }

  openModal(modalDial: any, post : any) {
    this.collectionInput = '';
    this.postToChange = post;
    this.modalService.open(modalDial, { ariaLabelledBy: 'modal-basic-title' }).result.then((result: any) => {
      console.log(`Closed`);
    }, (reason: any) => {
      console.log(`Dismissed`);
    });
  }

  changeCollection(): void {
    if(this.collectionInput.trim().length == 0){
      alert("Morate uneti ime kolekcije")
    }
    let id = localStorage.getItem("id") as string;
    let collection : PostCollection = {name : this.collectionInput};
    let changePost : FavPost = {postid : this.postToChange.id, userid : id, ownerid  : this.postToChange.userid, collection : collection};
    this.postsService.changeCollection(changePost).subscribe(items => {
      this.collections = [];
      this.posts = [];
      this.postsToShow = [];
      this.allArchivedPosts = [];
      this.initData();
      this.modalService.dismissAll();
    });
  }

  imageClick = (post: any) => {
    this.router.navigate(['single-post/' + post.userid + '/' + post.id]);
  };

}
