import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { UploadComponent } from './upload/upload.component';

const routes: Routes = [
  {path: '', redirectTo: '/upload', pathMatch: 'full'},
  {path: 'login', component: LoginComponent}, 
  {path: 'register', component: RegisterComponent},
  {path: 'upload', component: UploadComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
