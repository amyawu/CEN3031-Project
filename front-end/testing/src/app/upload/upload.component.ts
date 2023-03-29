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

  onFileSelected(event) {
    console.log(event)
    this.selectedFile = <File>event.target.files[0]
  }

  onUpload() {
    const fdata = new FormData();
    let test_name = "/C:/Users/manuc/Downloads/" + this.selectedFile.name 
    //fdata.append('file', this.selectedFile, this.selectedFile.name);
    fdata.append('file', this.selectedFile, test_name);
    console.log("DEBUGGING MODE")
    console.log(fdata.getAll)
    console.log(fdata)

    console.log(this.selectedFile.name);
    console.log(this.selectedFile.webkitRelativePath);

    console.log(test_name);

    this.http.post('http://localhost:8000/file', fdata)
    .subscribe(res => {
      console.log(res);
    });
  }
}
