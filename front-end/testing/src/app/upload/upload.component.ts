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

    if (!this.selectedFile) {
      console.error('No file selected');
      return;
    }
  
    const fdata = new FormData();
    fdata.append('file', this.selectedFile, this.selectedFile.name);
  

    this.http.post('http://localhost:8000/file', fdata)
      .subscribe(res => {
        console.log(res);
        this.uploadUserData = res;
      });



    // const fdata = new FormData();
    // let test_name = "/C:/Users/manuc/Downloads/" + this.selectedFile.name 
    // //fdata.append('file', this.selectedFile, this.selectedFile.name);
    // fdata.append('file', this.selectedFile, test_name);
    // console.log("DEBUGGING MODE")
    // console.log(fdata.getAll)
    // console.log(fdata)

    // console.log(this.selectedFile.name);
    // console.log(this.selectedFile.webkitRelativePath);

    // console.log(test_name);

    // this.http.post('http://localhost:8000/file', fdata)
    // .subscribe(res => {
    //   console.log(res);
    // });
  }
}
