import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-recents',
  templateUrl: './recents.component.html',
  styleUrls: ['./recents.component.css']
})
export class RecentsComponent {
  recentsUserData:any = {};

  constructor(private _auth: AuthService) {}

  recentsUser() {
    
    console.log(this.recentsUserData)

    this._auth.recentsUser(this.recentsUserData)
    .subscribe(
      (      res: any) => console.log(res),
      (      err: any) => console.log(err)
    )
  }
}
