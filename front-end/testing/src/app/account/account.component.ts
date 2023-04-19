import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-profile',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})
export class AccountComponent {
  accountUserData:any = {};

  constructor(private _auth: AuthService, private http: HttpClient) {
    this.updateProfile();
  }

  get authService() {
    return this._auth;
  }

  updateProfile() {
    let token = localStorage.getItem('token');
    let tstring = '{"token": ' + '"' + token + '"}';
    let jsonobj = JSON.parse(tstring);
    

    this.http.put<any>('http://localhost:8000/users/profile', jsonobj).subscribe(
      res => {
        console.log(res)
        this.accountUserData = res;
      },
      err => {
        console.log(err);
      }
    );
  }

  
}
