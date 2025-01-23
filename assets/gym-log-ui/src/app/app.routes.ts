import { Routes } from '@angular/router';
import { DashboardComponent } from './shared/components/dashboard/dashboard.component';
import { BodyComponent } from './features/body/body.component';
import { WorkoutComponent } from './features/workout/workout.component';
import { UserComponent } from './features/user/user.component';
import { GoalComponent } from './features/goal/goal.component';
import { SettingsComponent } from './shared/components/settings/settings.component';
import { PageNotFoundComponent } from './shared/components/page-not-found/page-not-found.component';
import { WelcomeComponent } from './shared/components/welcome/welcome.component';
import { LoginComponent } from './shared/components/login/login.component';
import { NewUserComponent } from './features/user/new-user/new-user.component';

export const routes: Routes = [
    { path: "", component: DashboardComponent, title: "Home"},
    { path: "home", component: DashboardComponent, title: "Home"},
    { path: "settings", component: SettingsComponent, title: "Settings"},
    { path: "user", component: UserComponent, title: "User"},
    { path: "workouts", component: WorkoutComponent, title: "Workouts"},
    { path: "goals", component: GoalComponent, title: "Goals"},
    { path: "body", component: BodyComponent, title: "Body"},
    { path: "welcome", component: WelcomeComponent, title: "Welcome"},
    { path: "login", component: LoginComponent, title: "Login"},
    { path: "user/new", component: NewUserComponent, title: "Create Profile"},
    { path: "api/register/google", component: DashboardComponent, title: "New User"},

    { path: "**", component: PageNotFoundComponent, title: "Page Not Found"},
];
