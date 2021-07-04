import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { Router, RouterModule, Routes } from '@angular/router';
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
import { FormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { VerificationRequestsComponent } from './verification-requests/verification-requests.component';
import { PostsComponent } from './home-page/posts/posts.component';
import { TokenInterceptor } from './interceptor/TokenInterceptor';
import { LikedDislikedPostsComponent } from './home-page/liked-disliked-posts/liked-disliked-posts.component';
import { ReportedPostsComponent } from './reported-posts/reported-posts.component';
import { AgentRequestsComponent } from './agent-requests/agent-requests.component';
import { SinglePostComponent } from './single-post/single-post.component';
import { SavedPostsComponent } from './saved-posts/saved-posts.component';
import { SingleCampaignComponent } from './single-campaign/single-campaign.component';

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
  {path: 'verificationRequests', component: VerificationRequestsComponent},
  {path: 'agentRequests', component: AgentRequestsComponent},
  {path: 'homePage/tag/:tag', component: HomePageComponent },
  {path: 'single-post/:userid/:postid', component: SinglePostComponent },
  {path: 'single-campaign/:userid/:campaignid', component: SingleCampaignComponent },
  {path: 'reportedPosts', component:ReportedPostsComponent},
  {path: 'homePage/location/:location', component: HomePageComponent },
  {path: 'favourites', component: SavedPostsComponent },
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
    VerificationRequestsComponent,
    PostsComponent,
    LikedDislikedPostsComponent,
    ReportedPostsComponent,
    AgentRequestsComponent,
    SinglePostComponent,
    SavedPostsComponent,
    SingleCampaignComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(appRoutes, { onSameUrlNavigation: 'reload' }),
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    NgbModule,
  ],
  exports: [
    RouterModule
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: TokenInterceptor,
      multi: true,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
