import { Component } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  registerUserData:any = {};

  constructor(private _auth: AuthService) {}

  registerUser() {
    
    console.log(this.registerUserData)

    this._auth.registerUser(this.registerUserData)
    .subscribe(
      (      res: any) => console.log(res),
      (      err: any) => console.log(err)
    )
  }
}
