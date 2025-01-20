import { Routes } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { BodyComponent } from './body/body.component';
import { WorkoutComponent } from './workout/workout.component';
import { UserComponent } from './user/user.component';
import { GoalComponent } from './goal/goal.component';
import { SettingComponent } from './setting/setting.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { WelcomeComponent } from './welcome/welcome.component';
import { LoginComponent } from './login/login.component';
import { NewUserComponent } from './new-user/new-user.component';

export const routes: Routes = [
    { path: "", component: DashboardComponent, title: "Home"},
    { path: "home", component: DashboardComponent, title: "Home"},
    { path: "settings", component: SettingComponent, title: "Settings"},
    { path: "user", component: UserComponent, title: "User"},
    { path: "workouts", component: WorkoutComponent, title: "Workouts"},
    { path: "goals", component: GoalComponent, title: "Goals"},
    { path: "body", component: BodyComponent, title: "Body"},
    { path: "welcome", component: WelcomeComponent, title: "Welcome"},
    { path: "login", component: LoginComponent, title: "Login"},
    { path: "user/new", component: NewUserComponent, title: "New User"},

    { path: "**", component: PageNotFoundComponent, title: "Page Not Found"},
];
