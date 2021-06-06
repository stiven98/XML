import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {RouterModule, Routes} from '@angular/router';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { RegisterComponent } from './register/register.component';
import { LoginComponent } from './login/login.component';
import { ProfilePageComponent } from './profile-page/profile-page.component';
import { EditProfileComponent } from './edit-profile/edit-profile.component';
import { DirectMessagesComponent } from './direct-messages/direct-messages.component';
import { NotificationsComponent } from './notifications/notifications.component';
import { HomePageComponent } from './home-page/home-page.component';
import {FormsModule} from '@angular/forms';
import {HttpClientModule} from '@angular/common/http';

const appRoutes: Routes = [
  {path: 'registration', component: RegisterComponent },
  {path: 'login', component: LoginComponent },
  {path: 'profile', component: ProfilePageComponent },
  {path: 'editProfile', component: EditProfileComponent },
  {path: 'directMessages', component: DirectMessagesComponent },
  {path: 'notifications', component: NotificationsComponent },
  {path: 'homePage', component: HomePageComponent },
  { path: '**', redirectTo: '/404'}
];

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    RegisterComponent,
    LoginComponent,
    ProfilePageComponent,
    EditProfileComponent,
    DirectMessagesComponent,
    NotificationsComponent,
    HomePageComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(appRoutes),
    AppRoutingModule,
    FormsModule,
    HttpClientModule
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
