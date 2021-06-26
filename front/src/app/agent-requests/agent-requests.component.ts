import { Component, OnInit } from '@angular/core';
import { AgetnRegistrationRequest } from '../model/User.model';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-agent-requests',
  templateUrl: './agent-requests.component.html',
  styleUrls: ['./agent-requests.component.css']
})
export class AgentRequestsComponent implements OnInit {

  agentRegistrationRequests: any[] | undefined;
  constructor(private userService: UserService) { }

  ngOnInit(): void {
   this.initData();
  }
  initData = () => {
    this.userService.getAllAgentRequests().subscribe(response => {
      this.agentRegistrationRequests = response as any;
    });
  }

  accept = (request:any) => {
    this.userService.createAgent(request).subscribe(res => {
      res;
    });
    this.initData();
  }
  decline = (request:any) => {
    this.userService.declineAgent(request).subscribe(res => {res});
    this.initData();
  }

}
