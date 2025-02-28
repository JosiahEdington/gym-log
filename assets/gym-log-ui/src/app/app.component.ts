import { Component } from '@angular/core';
import { RouterOutlet, Router, TitleStrategy } from '@angular/router';
import { NavbarComponent } from "./shared/components/navbar/navbar.component";
import { Title } from '@angular/platform-browser';
import { NgIf } from '@angular/common';
import { SocialLoginModule, SocialAuthServiceConfig } from '@abacritt/angularx-social-login'
import { GoogleLoginProvider } from '@abacritt/angularx-social-login'

@Component({
    selector: 'app-root',
    imports: [
        RouterOutlet,
        NavbarComponent,
        NgIf,
        SocialLoginModule,
    ],
    templateUrl: './app.component.html',
    styleUrl: './app.component.css'
})
export class AppComponent {
  title: string = 'gym-log-ui'
  hiddenPages: Array<string> = [
    // TODO: Set up a route in the Welcome page to log in.
    // 'Welcome',
    'Page Not Found',
    'Login',
    'Create Profile',
  ]; 
  
  constructor(private router: Router, private titles: Title) {}

  hasNavbar(): boolean {
  if (this.hiddenPages.includes(this.titles.getTitle())) {
    return false
  }
  return true
  }
}
