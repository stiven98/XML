import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { DeletePost } from '../model/DeletePost';
import { ReportReq } from '../model/ReportReq';
import { AuthService } from '../service/authorization/auth.service';
import { PostsService } from '../service/posts.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-single-campaign',
  templateUrl: './single-campaign.component.html',
  styleUrls: ['./single-campaign.component.css'],
})
export class SingleCampaignComponent implements OnInit {
  public campaign: any;
  public userData: Map<string, any> = new Map<string, any>();
  public commentData: Map<string, string> = new Map<string, string>();
  public reportReq: ReportReq = new ReportReq();
  public currentUser: any;
  public commentText: string = '';
  public userid: string = '';
  public campaignid: string = '';
  public allowDelete: boolean = false;
  public allowChange: boolean = false;
  public isMultiple: boolean = false;
  public startday: any;
  public endday: any;
  public isInfluencer: boolean = false;

  constructor(
    public postsService: PostsService,
    public userService: UserService,
    public authService: AuthService,
    public router: Router,
    public route: ActivatedRoute,
    private modalService: NgbModal
  ) {}

  ngOnInit(): void {
    this.campaignid = this.route.snapshot.params[`campaignid`];
    this.userid = this.route.snapshot.params[`userid`];
    this.userService.getUserById(this.userid).subscribe((us: any) => {
      console.log(us);
      console.log('rezzzzzzz');
      if (us.isVerified) {
        this.isInfluencer = true;
      }
      console.log(this.isInfluencer);
      this.initData();
    });
  }
  initData = () => {
    if (localStorage.getItem('id') === this.userid) {
      this.allowChange = true;
      this.allowDelete = true;
    }
    let id = localStorage.getItem('id');

    if (!this.isInfluencer) {
      console.log('usao u if');
      this.postsService
        .getCampaignByIds(this.userid, this.campaignid)
        .subscribe((res) => {
          this.campaign = res as any;
          this.startday = this.campaign.startday.split('T')[0];
          this.endday = this.campaign.endday.split('T')[0];
          this.isMultiple = this.campaign.ismultiple;
          this.userService.getUsersById(id).subscribe((response) => {
            this.currentUser = response;
          });

          if (this.campaign) {
            console.log('usao u zadnji if');
            if (!this.userData.get(this.campaign.userid)) {
              this.userService
                .getUsersById(this.campaign.userid)
                .subscribe((res) => {
                  this.userData.set(this.campaign.userid, res);
                  console.log(this.campaign);
                });
            }
          }
        });
    } else {
      this.postsService
        .getCampaignByInfluencerIds(this.userid, this.campaignid)
        .subscribe((result) => {
          this.campaign = result as any;
          this.startday = this.campaign.startday.split('T')[0];
          this.endday = this.campaign.endday.split('T')[0];
          this.isMultiple = this.campaign.ismultiple;
          this.userService.getUsersById(id).subscribe((response) => {
            this.currentUser = response;
          });

          if (this.campaign) {
            console.log('usao u zadnji if');
            if (!this.userData.get(this.campaign.userid)) {
              this.userService
                .getUsersById(this.campaign.userid)
                .subscribe((res) => {
                  this.userData.set(this.campaign.userid, res);
                  console.log(this.campaign);
                });
            }
          }
        });
    }
  };

  onDeleteClick = () => {
    let deletePost = new DeletePost();
    deletePost.ownerId = this.userid;
    deletePost.postId = this.campaign.id;
    this.postsService.deleteCampaign(deletePost).subscribe((res) => {
      if (res) {
        alert('Kampanja uspesno obrisana');
        this.router.navigate(['homePage']);
      } else {
        alert('Došlo je do greške');
      }
    });
  };

  openModal(modalDial: any) {
    this.modalService
      .open(modalDial, { ariaLabelledBy: 'modal-basic-title' })
      .result.then(
        (result: any) => {
          console.log(`Closed`);
        },
        (reason: any) => {
          console.log(`Dismissed`);
        }
      );
  }

  onUpdateClick = () => {
    this.campaign.startday = this.startday + 'T01:00:00+01:00';
    this.campaign.endday = this.endday + 'T01:00:00+01:00';
    this.postsService.updateCampaign(this.campaign).subscribe((res) => {
      alert('Kampanja uspešno izmenjena!');
      this.router.navigate(['homePage']);
    });
  };

  getCurrentUserImage = (): string => {
    return this.currentUser.system_user.picturePath;
  };

  getUserImage = (comment: any): string => {
    var user = this.userData.get(comment.userid);
    return user.system_user.picturePath;
  };

  getUsername = (id: string): string => {
    var user = this.userData.get(id);
    if (user) return user.system_user.username;
    else return '';
  };

  onUsernameClick = (event: Event, id: string) => {
    event.preventDefault();
    var user = this.userData.get(id);
    if (user) this.router.navigate(['/profile', id]);
  };

  isImage = (name: string): boolean => {
    let imgFormats = ['jpg', 'jpeg', 'gif', 'png', 'tiff', 'bmp'];
    let flag = false;
    for (let i = 0; i < imgFormats.length; i++) {
      if (name.includes(imgFormats[i])) {
        flag = true;
        break;
      }
    }
    return flag;
  };

  onViewMoreClick = () => {
    alert('Aca');
    let link = 'www.google.com';
    window.location.href = link;
  };
  likeClick = (post: any) => {
    let postId = post.id;
    let ownerId = post.userid;
    let id = localStorage.getItem('id');
    if (id) {
      this.postsService
        .likePost({ userid: id, postid: postId, ownerid: ownerId })
        .subscribe((res) => {
          this.initData();
        });
    } else {
      alert('Morate biti ulogovani da bi lajkovali objavu');
    }
  };

  dislikeClick = (post: any) => {
    let postId = post.id;
    let ownerId = post.userid;
    let id = localStorage.getItem('id');
    if (id) {
      this.postsService
        .dislikePost({ userid: id, postid: postId, ownerid: ownerId })
        .subscribe((res) => {
          this.initData();
        });
    } else {
      alert('Morate biti ulogovani da bi dislajkovali objavu');
    }
  };
}
