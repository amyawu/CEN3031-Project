import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-display',
  templateUrl: './display.component.html',
  styleUrls: ['./display.component.css']
})
export class DisplayComponent {
  displayUserData: any = {};
  imageUrls: any = [];

  constructor(private _auth: AuthService, private http: HttpClient) {
    this.displayUser();
  }

  get authService() {
    return this._auth;
  }

  displayUser() {
    let token = localStorage.getItem('token');
    let tstring = '{"token": ' + '"' + token + '"}';
    let jsonobj = JSON.parse(tstring);

    this.http.post<any>('http://localhost:8000/users/all_images', jsonobj).subscribe(
      res => {
        console.log(res);
        this.displayUserData = res;
        this.imageUrls = res.urls;
      },
      err => {
        console.log(err);
      }
    );
  }
}
