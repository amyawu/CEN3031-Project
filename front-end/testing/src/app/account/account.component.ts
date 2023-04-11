import { Component } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-profile',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent {
  accountUserData:any = {};

  constructor(private _auth: AuthService) {}

  accountUser() {
    
    console.log(this.accountUserData)

    this._auth.accountUser(this.accountUserData)
    .subscribe(
      (      res: any) => console.log(res),
      (      err: any) => console.log(err)
    )
  }
}
