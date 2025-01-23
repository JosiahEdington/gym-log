import { UserSetting } from "./user-setting";
export class User {
    UserId?: number;
    FirstName?: string;
    LastName?: string;
    Username?: string;
    Password?: string;
    Picture?: string;
    Email?: string;
    Phone?: string;
    UserSettings?: Array<UserSetting>
}
