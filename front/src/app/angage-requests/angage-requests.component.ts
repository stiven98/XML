import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AddInfluencer } from '../model/AddInfluencer';
import { DeleteReq } from '../model/DeleteReq';
import { PostsService } from '../service/posts.service';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-angage-requests',
  templateUrl: './angage-requests.component.html',
  styleUrls: ['./angage-requests.component.css'],
})
export class AngageRequestsComponent implements OnInit {
  public angageRequests: any[] = [];
  public agent: any;
  public deleteReq: DeleteReq = new DeleteReq();
  public addInfluencer: AddInfluencer = new AddInfluencer();
  constructor(private postService: PostsService, private router: Router) {}

  ngOnInit(): void {
    this.initData();
  }
  initData = () => {
    this.postService
      .getAngageRequests(localStorage.getItem('id') as string)
      .subscribe((res: any) => {
        this.angageRequests = res;
        console.log(this.angageRequests);
      });
  };

  onViewCampaignClick = (agentId: string, campaignId: string) => {
    this.router.navigate(['/single-campaign/' + agentId + '/' + campaignId]);
  };
  onDeleteClick = (influencerId: string, id: string) => {
    this.deleteReq.ownerId = influencerId;
    this.deleteReq.campaignReqId = id;
    this.postService.deleteCampaignReq(this.deleteReq).subscribe((res) => {
      alert('Zahtev uspesno odbijen!');
      this.router.navigate(['/homePage']);
    });
  };
  onAcceptClick = (
    influencerId: string,
    campaignId: string,
    ownerId: string,
    id: string
  ) => {
    this.addInfluencer.campaignId = campaignId;
    this.addInfluencer.influencerId = influencerId;
    this.addInfluencer.ownerId = ownerId;
    this.addInfluencer.id = id;
    this.postService.addInfluencer(this.addInfluencer).subscribe((res: any) => {
      alert('Zahtev uspesno prihvacen!');
      this.router.navigate(['/homePage']);
    });
  };
}
