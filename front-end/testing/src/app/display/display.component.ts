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
    
    console.log(this.displayUserData)

    this._auth.displayUser(this.displayUserData)
    .subscribe(
      (      res: any) => console.log(res),
      (      err: any) => console.log(err)
    )
  }
}
