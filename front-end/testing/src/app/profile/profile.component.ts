import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  profileUserData:any = {};

  constructor(private _auth: AuthService, private _router: Router) {}

  get authService() {
    return this._auth;
  }

  profileUser() {

    this.profileUserData.age = parseInt(this.profileUserData.age, 10);
  
    let token = localStorage.getItem('token');
    console.log(typeof token)
    let tstring = '{"token": ' + '"' + token + '"}';
    console.log(tstring)
    console.log(this.profileUserData)
    let jsonobj = JSON.parse(tstring)
    console.log(jsonobj)

    let full = Object.assign({}, jsonobj, this.profileUserData);


    this._auth.submitProfile(full)
    .subscribe(
      (      res: any) => {
        console.log(res)
        this._router.navigate(['/account'])},
      (      err: any) => console.log(err)
    )
  }
}
