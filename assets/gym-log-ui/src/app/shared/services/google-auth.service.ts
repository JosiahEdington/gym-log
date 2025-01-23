import { Injectable, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { SocialAuthService } from '@abacritt/angularx-social-login';
import { GoogleLoginProvider } from '@abacritt/angularx-social-login';
import { SocialUser } from '@abacritt/angularx-social-login';

@Injectable({
  providedIn: 'root',
})
export class GoogleAuthService implements OnInit {
  private accessToken = '';
  socialUser?: SocialUser;
  loggedIn: boolean = false;

  constructor(private authService: SocialAuthService, private httpClient: HttpClient) { }

  ngOnInit() {
    this.authService.authState.subscribe((user) => {
      this.socialUser = user;
      this.loggedIn = (user != null);
    });
  }

  getAccessToken(): void {
    this.authService.getAccessToken(GoogleLoginProvider.PROVIDER_ID).then(
          accessToken => this.accessToken = accessToken
    );
  }

  getGoogleCalendarData(): void {
    if (!this.accessToken) return;

    this.httpClient
        .get('https://www.googleapis.com/calendar/v3/calendars/primary/events', {
          headers: { Authorization: `Bearer ${this.accessToken}` },
        })
        .subscribe((events) => {
          alert("Look at your console");
          console.log('events', events);
      });
  }
  refreshToken(): void {
    this.authService.refreshAccessToken(GoogleLoginProvider.PROVIDER_ID);
  }

}
