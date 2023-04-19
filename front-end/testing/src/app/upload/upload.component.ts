import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})

export class UploadComponent {
  uploadUserData:any = {};
  selectedFile:File = new File([], '');

  constructor(private http: HttpClient, private _auth: AuthService) {}

  get authService() {
    return this._auth;
  }

  onFileSelected(event) {
    console.log(event)
    this.selectedFile = <File>event.target.files[0]
  }

  onUpload() {

    if (!this.selectedFile || !this.selectedFile.name) {
      alert('No file selected!');
      return;
    }
    else{
  
    const fdata = new FormData();
    fdata.append('file', this.selectedFile, this.selectedFile.name);
  
    const headers = {};
  
    // Check if the user is authenticated
    if (this.authService.loggedIn()) {
      // If the user is authenticated, set the Authorization header
      const token = localStorage.getItem('token')
      headers['Authorization'] = `Bearer ` + token;
      this.http.put('http://localhost:8000/users/image', fdata, { headers })
        .subscribe(res => {
          console.log(res);
          this.uploadUserData = res;
        });
    } else {
      // If the user is not authenticated, send the image to a different endpoint
      this.http.post('http://localhost:8000/file', fdata)
        .subscribe(res => {
          console.log(res);
          this.uploadUserData = res;
        });
  }
  }
}
}
