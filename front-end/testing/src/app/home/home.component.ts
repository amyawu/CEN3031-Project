import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  homeUserData:any = {};

  constructor(private _auth: AuthService) {}

  homeUser() {
    console.log(this.homeUserData)

    this._auth.homeUser(this.homeUserData)
    .subscribe(
      (      res: any) => console.log(res),
      (      err: any) => console.log(err)
    )
  }
}
