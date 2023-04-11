import { Component } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  profileUserData:any = {};

  constructor(private _auth: AuthService) {}

  profileUser() {
    
    console.log(this.profileUserData)

    this._auth.profileUser(this.profileUserData)
    .subscribe(
      (      res: any) => console.log(res),
      (      err: any) => console.log(err)
    )
  }
}
