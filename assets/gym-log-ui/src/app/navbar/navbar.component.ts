import { CommonModule } from '@angular/common';
import { trigger, style, animate, transition } from '@angular/animations';
import { Component } from '@angular/core';
import { RouterModule, Router } from '@angular/router';

@Component({
  selector: 'app-navbar',
  standalone: true,
  animations: [
    trigger('enterAnimation',[
                transition(':enter', [
                  style({transform: 'opacity-0 scale-95'}),
                  animate('500ms', style({transform: 'opacity-100 scale-100'}))
                ]),
                transition(':leave', [
                  style({transform: 'opacity-100 scale-100'}),
                  animate('500ms', style({transform: 'opacity-0 scale-95'}))
                ])
              ]
            )
          ],
  imports: [
    RouterModule,
    CommonModule,
],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css'
})
export class NavbarComponent {
  showDropdown: boolean = false;

  constructor(private router: Router) {}

  isActive(route: string): boolean {
    return this.router.url === route;
  }

  toggleDropdown() {
    this.showDropdown = !this.showDropdown;
  }
}
