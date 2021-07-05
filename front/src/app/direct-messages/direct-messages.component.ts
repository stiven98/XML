import { Component, OnInit } from '@angular/core';
import {UserService} from '../service/user.service';
import {Event} from '@angular/router';
import {AuthService} from '../service/authorization/auth.service';
import {MessagesService} from '../service/messages.service';
import {Message} from '../model/Message';
import {window} from 'rxjs/operators';

@Component({
  selector: 'app-direct-messages',
  templateUrl: './direct-messages.component.html',
  styleUrls: ['./direct-messages.component.css']
})
export class DirectMessagesComponent implements OnInit {

  public users: any[] = [];
  public messages: any[] = [];
  public selectedUserID = '';
  public selectedUserImage = '';
  public textMessage = '';
  public myId: string | null = '';
  public message: Message | undefined;



  constructor(private userService: UserService, private authService: AuthService, private messagesService: MessagesService) { }

  ngOnInit(): void {

    this.myId = this.authService.getId();


    this.userService.getAllUsers().subscribe(response => {
      this.users = response as any [];
      console.log(this.users);
    });

  }

  onClickUser = (id: string, path: string) =>  {
    this.selectedUserID = id;
    this.selectedUserImage = path;
    console.log(this.selectedUserImage);
    this.messagesService.getConversation(this.authService.getId(), this.selectedUserID).subscribe(response => {
      console.log(response);
      // @ts-ignore
      this.messages = response.messages as [];
    });
  }



  send = () => {
    console.log(this.textMessage);
    this.message = new Message();
    this.message.sender = this.myId;
    this.message.receiver = this.selectedUserID;
    this.message.content = this.textMessage;

    this.messagesService.sendTextMessage(this.myId, this.selectedUserID, this.message).subscribe(response => {
      console.log(response);
      // @ts-ignore

      this.messages = response.messages as [];
      this.textMessage = '';
    });
  }

  uploadImage = (event: any) => {
    event.preventDefault();
    const formData = new FormData();
    formData.append('files', event.target.files[0]);
    this.messagesService.upload(formData).subscribe(response => {
      console.log(response);
    });
  }
}
