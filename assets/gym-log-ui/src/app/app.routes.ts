import { Routes } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { BodyComponent } from './body/body.component';
import { WorkoutComponent } from './workout/workout.component';
import { UserComponent } from './user/user.component';
import { GoalComponent } from './goal/goal.component';
import { SettingComponent } from './setting/setting.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';

export const routes: Routes = [
    { path: "", component: DashboardComponent},
    { path: "settings", component: SettingComponent},
    { path: "user", component: UserComponent},
    { path: "workouts", component: WorkoutComponent},
    { path: "goals", component: GoalComponent},
    { path: "body", component: BodyComponent},
    { path: "**", component: PageNotFoundComponent},
];
