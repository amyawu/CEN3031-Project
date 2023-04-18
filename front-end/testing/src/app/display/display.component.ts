import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-display',
  templateUrl: './display.component.html',
  styleUrls: ['./display.component.css']
})
export class DisplayComponent {
  displayUserData:any = {};

  constructor(private _auth: AuthService) {}

  displayUser() {
    let token = localStorage.getItem('token');
    let tstring = '{"token": ' + '"' + token + '"}';
    let jsonobj = JSON.parse(tstring);

    const imageURL = jsonobj;
    console.log("this is the image url:" +jsonobj);
  }
}
