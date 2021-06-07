import { Component, OnInit } from '@angular/core';
import { UserEdit } from '../model/EditUser';
import { UserService } from '../service/user.service';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit {
  edit: UserEdit = new UserEdit();
  confirm_password: string = '';
  new_password:string = '';
  date: Date = new Date('02/02/1998');
  constructor(private userService: UserService) { }

  ngOnInit(): void {
    let id:any = localStorage.getItem('id');
    this.userService.getUsersById(id).subscribe((res) => {
      this.edit = res as UserEdit;
    });
  }



  upDate = () => {
    this.validePssword()
    this.userService.editUser(this.edit).subscribe((res) => {console.log(res)})
  }

  validePssword = () => {
    if(this.new_password !== '' && this.confirm_password !== ''){
      if(this.new_password === this.confirm_password){
        this.edit.system_user.password = this.new_password
      }
    }
  }


}
