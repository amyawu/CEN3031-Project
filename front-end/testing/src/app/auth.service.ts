import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  //private _registerURL = "http://localhost:3000/api/register";
  private _registerURL = "http://localhost:8000/users";

  //private _loginURL = "http://localhost:3000/api/login";
  private _loginURL = "http://localhost:8000/users/login";

  private _profileURL = "http://localhost:8000/users/profile";

  private _homeURL = "http://localhost:8000/users/home";

  private _recentsURL = "http://localhost:8000/users/recents";

  private _accountURL = "http://localhost:8000/users/account";

  private _displayURL = "http://localhost:8000/users/display";

  constructor(private http: HttpClient, private _router: Router) { }

  registerUser(user: any) {
    return this.http.post<any>(this._registerURL, user)
  }

  loginUser(user : any) {
    return this.http.post<any>(this._loginURL, user)
  }

  profileUser(user : any) {
    return this.http.post<any>(this._profileURL, user)
  }

  homeUser(user : any) {
    return this.http.post<any>(this._homeURL, user)
  }

  recentsUser(user : any) {
    return this.http.post<any>(this._recentsURL, user)
  }

  accountUser(user : any) {
    return this.http.post<any>(this._accountURL, user)
  }

  displayUser(user : any) {
    return this.http.post<any>(this._displayURL, user)
  }

  loggedIn() {
    return localStorage.getItem('token') !== undefined && localStorage.getItem('token') !== null;
  }

  logoutUser() {
    localStorage.removeItem('token')
    this._router.navigate(['/login'])
  }

  submitProfile(user: any) {
    return this.http.put<any>(this._profileURL, user)
  }
}
