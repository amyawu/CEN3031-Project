import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})
export class UploadComponent {
  selectedFile:File = new File([], '');

  constructor(private http: HttpClient) {}

  onFileSelected(event: any) {
    console.log(event)
    this.selectedFile = <File>event.target.files[0]
  }

  onUpload() {
    const fdata = new FormData();
    fdata.append('image', this.selectedFile, this.selectedFile.name);
    this.http.post('http://localhost:8000/file', fdata)
    .subscribe(res => {
      console.log(res);
    });
  }
}
