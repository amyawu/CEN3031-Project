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
  displayUserData:any = {};

  constructor(private _auth: AuthService, private http: HttpClient) {
    this.displayUser();
  }

  displayUser() {
    let token = localStorage.getItem('token');
    let tstring = '{"token": ' + '"' + token + '"}';
    let jsonobj = JSON.parse(tstring);

    this.http.put<any>('http://localhost:8000/users/display', jsonobj).subscribe(
      res => {
        console.log(res)
        this.displayUserData = res;
      },
      err => {
        console.log(err);
      }
    );
  }


}
