import { Component } from '@angular/core';
import { ReactiveFormsModule, FormGroup, FormControl } from '@angular/forms';

@Component({
  selector: 'app-new-user',
  standalone: true,
  imports: [
    ReactiveFormsModule,
  ],
  templateUrl: './new-user.component.html',
  styleUrl: './new-user.component.css'
})
export class NewUserComponent {
  newUserForm = new FormGroup({
    username: new FormControl(''),
    password: new FormControl(''),
    // picture: new FormControl(''),
    firstName: new FormControl(''),
    lastName: new FormControl(''),
    email: new FormControl(''),
    phone: new FormControl(''),
    workoutReminders: new FormControl(false),
    workoutTracking: new FormControl(false),
    workoutAccountability: new FormControl(false),
    goalReminders: new FormControl(false),
    goalTracking: new FormControl(false),
    goalProgress: new FormControl(false),
    goalAccountability: new FormControl(false),
    pushNotifications: new FormControl("Phone"),
  })

  onSubmit() {
    console.log("Submitting form for new user...")
    console.log(this.newUserForm.value)
  }

}
