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
import {HttpClientModule, HTTP_INTERCEPTORS} from '@angular/common/http';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { PostsComponent } from './home-page/posts/posts.component';
import { TokenInterceptor } from './interceptor/TokenInterceptor';
import { LikedDislikedPostsComponent } from './liked-disliked-posts/liked-disliked-posts.component';

const appRoutes: Routes = [
  {path: 'registration', component: RegisterComponent },
  {path: 'login', component: LoginComponent },
  {path: 'profile/:id', component: ProfilePageComponent },
  {path: 'editProfile', component: EditProfileComponent },
  {path: 'directMessages', component: DirectMessagesComponent },
  {path: 'notifications', component: NotificationsComponent },
  {path: 'liked', component: LikedDislikedPostsComponent},
  {path: 'disliked', component: LikedDislikedPostsComponent},
  {path: 'homePage', component: HomePageComponent },
  {path: 'homePage/tag/:tag', component: HomePageComponent },
  {path: 'homePage/location/:location', component: HomePageComponent },
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
    PostsComponent,
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(appRoutes),
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    NgbModule
    ],
    providers: [
      {
      provide: HTTP_INTERCEPTORS,
      useClass: TokenInterceptor,
      multi: true
      }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
