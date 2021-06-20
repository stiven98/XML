import { Component, OnInit } from '@angular/core';
import { UserEdit } from '../model/EditUser';
import { UserService } from '../service/user.service';
import {AuthService} from '../service/authorization/auth.service';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {Post, PostItem} from '../model/Post.model';
import {VerificationRequestService} from '../service/verificationRequest.service';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit {

  edit: UserEdit = new UserEdit();
  confirm_password: string = '';
  new_password: string = '';
  date: Date = new Date('02/02/1998');

  document: any = undefined;
  changeProfilePictureFlag = true;
  url = '';

  constructor(private userService: UserService,
              public authService: AuthService, private modalService: NgbModal,
              private verificationRequestService: VerificationRequestService) { }



  ngOnInit(): void {
    let id: any = localStorage.getItem('id');
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
    this.validePssword();
    this.edit.system_user.dateOfBirth = this.edit.system_user.dateOfBirth +'T01:00:00+01:00'
    this.userService.editUser(this.edit)
    .subscribe((res) => {
      this.edit.system_user.dateOfBirth = this.edit.system_user.dateOfBirth.split("T")[0];
       alert("Uspesno ste izmenili podatke")}
       , error => {if(error.status === 400) alert("Podaci nisu validni")});

  }

  validePssword = () => {
    if (this.new_password !== '' && this.confirm_password !== ''){
      if (this.new_password === this.confirm_password){
        this.edit.system_user.password = this.new_password;
      }
    }
  }

  openModal = (modalDial: any) => {
    this.document = undefined;
    this.modalService.open(modalDial, { ariaLabelledBy: 'modal-basic-title' }).result.then((result: any) => {
      console.log(`Closed`);
    }, (reason: any) => {
      console.log(`Dismissed`);
    });
  }

  verify = () => {
    if (this.document !== undefined && this.authService.getId() !== null) {
      const formData = new FormData();

      formData.append('document', this.document);


      this.verificationRequestService.verify(this.authService.getId(), formData).subscribe((response) => {
        console.log(response);
        alert('Zahtjev kreiran!');
      });


    }



  }

  onChangeVerify = (event: any): void => {
    this.document = event.target.files[0];
    console.log(this.document);
  }

  changePictureFlag = () =>{
    this.changeProfilePictureFlag = !this.changeProfilePictureFlag
  }

  changeProfilePicture = () => {
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
      this.url = event.target.files[0];
  }

}
