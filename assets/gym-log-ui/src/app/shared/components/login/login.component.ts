import { Component} from '@angular/core';
import { Router } from '@angular/router';
import { GoogleLoginProvider, GoogleSigninButtonModule, SocialAuthService } from '@abacritt/angularx-social-login';
@Component({
  selector: 'app-login',
  standalone: true,
  imports: [
    GoogleSigninButtonModule,
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  constructor(private authService: SocialAuthService, private router: Router) {}

  // ngOnInit(): void {
  //   this.authService.authState.subscribe((user) => {
  //     console.log("Login successful: ", user)
  //     this.router.navigate(['/']);
  //   });
  // }
  refreshToken(): void {
    this.authService.refreshAuthToken(GoogleLoginProvider.PROVIDER_ID);
  }

  signInWithgoogle(): void {
    this.authService.authState.subscribe((user) => {
      // Do more user object stuff here
      console.log("Login successful: ", user)
      this.router.navigate(['/']);
    });
    console.log("Login failed.")
  }

}
