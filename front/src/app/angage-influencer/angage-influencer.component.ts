import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { CampaignRequest } from '../model/CampaignRequest';
import { AuthService } from '../service/authorization/auth.service';
import { PostsService } from '../service/posts.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-angage-influencer',
  templateUrl: './angage-influencer.component.html',
  styleUrls: ['./angage-influencer.component.css'],
})
export class AngageInfluencerComponent implements OnInit {
  public campaigns: any[] = [];
  public agentId: string = '';
  public currentUser: any;
  public currentAgent: any;
  public influencerId: string = '';
  public campaignRequest: CampaignRequest = new CampaignRequest();
  constructor(
    private route: ActivatedRoute,
    private postsService: PostsService,
    public userService: UserService,
    public authService: AuthService
  ) {
    this.agentId = route.snapshot.params[`agentid`];
    this.influencerId = route.snapshot.params[`influencerid`];
  }

  ngOnInit(): void {
    this.initData();
  }

  initData = () => {
    this.postsService
      .getCampaignsByUserId(this.agentId)
      .subscribe((result: any) => {
        this.campaigns = result as any[];
        this.userService.getUserById(this.agentId).subscribe((us) => {
          this.currentUser = us;
          this.userService.getUserById(this.influencerId).subscribe((inf) => {
            this.currentAgent = inf;
          })
        });
      });
  };
  getCurrentUserImage = (): string => {
    return this.currentUser.system_user.picturePath;
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

  createCampaignRequest = (id:string) => {
    this.campaignRequest.agentId = this.agentId;
    this.campaignRequest.influencerId = this.influencerId;
    this.campaignRequest.campaignId = id;
    this.postsService.createCampaignRequest(this.campaignRequest).subscribe((res: any) =>{ res});
  }
}
