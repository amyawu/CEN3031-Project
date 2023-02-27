import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})
export class UploadComponent {
  selectedFile = null;

  constructor(private http: HttpClient) {}

  onFileSelected(event: any) {
    console.log(event)
    this.selectedFile = event.target.files[0]
  }

  onUpload() {
    
  }
}
