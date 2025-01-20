import { Component } from '@angular/core';
import { RouterOutlet, Router, TitleStrategy } from '@angular/router';
import { NavbarComponent } from "./navbar/navbar.component";
import { Title } from '@angular/platform-browser';
import { NgIf } from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet, 
    NavbarComponent,
    NgIf,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  hiddenPages: Array<string> = [
    // TODO: Set up a route in the Welcome page to log in.
    //'Welcome',
    'Page Not Found',
    'Login',
    'New User',
  ]; 
  
  constructor(private router: Router, private title: Title) {}

  hasNavbar(): boolean {
  if (this.hiddenPages.includes(this.title.getTitle())) {
    return false
  }
  return true
  }
}
