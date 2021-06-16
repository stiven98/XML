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
  changeProfilePictureFlag = true;
  url = ''
  constructor(private userService: UserService) { }

  ngOnInit(): void {
    let id:any = localStorage.getItem('id');
    this.userService.getUsersById(id).subscribe((res) => {
      this.edit = res as UserEdit;
      this.edit.system_user.dateOfBirth = this.edit.system_user.dateOfBirth.split("T")[0];
    });
  }

  changeDate = (event: any) => {
    this.edit.system_user.dateOfBirth = event.target.value;
    console.log(this.edit.system_user.dateOfBirth);
  }


  upDate = () => {
    this.validePssword()
    this.edit.system_user.dateOfBirth = this.edit.system_user.dateOfBirth +'T01:00:00+01:00'
    this.userService.editUser(this.edit)
    .subscribe((res) => {  
      this.edit.system_user.dateOfBirth = this.edit.system_user.dateOfBirth.split("T")[0];
       alert("Uspesno ste izmenili podatke")}
       , error => {if(error.status === 400) alert("Podaci nisu validni")});
      
  }

  validePssword = () => {
    if(this.new_password !== '' && this.confirm_password !== ''){
      if(this.new_password === this.confirm_password){
        this.edit.system_user.password = this.new_password
      }
    }
  }

  changePictureFlag = () =>{
    this.changeProfilePictureFlag = !this.changeProfilePictureFlag
  }

  changeProfilePicture = () =>{
    let formData = new FormData();
    formData.append("files", this.url);

    this.userService.uploadProfilePicture(formData).subscribe((res) => {
      this.changeProfilePictureFlag = !this.changeProfilePictureFlag
      this.edit.system_user.picturePath = res as string;
      this.edit.system_user.dateOfBirth = this.edit.system_user.dateOfBirth +'T01:00:00+01:00'
      this.userService.editUser(this.edit).subscribe((res) => {  
        this.edit.system_user.dateOfBirth = this.edit.system_user.dateOfBirth.split("T")[0];
         alert("Uspesno ste izmenili sliku")});
    },
    error => {if(error.status == 500 || error.status ==400) alert("Neuspesna izmena slike")});
  }

  onChange(event: any): void {
      this.url =event.target.files[0];
  }

}
