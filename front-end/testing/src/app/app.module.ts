import { NgModule } from '@angular/core';
//import { NgModule, NO_ERRORS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import {MatFormFieldModule} from '@angular/material/form-field'; 
import {MatInputModule} from '@angular/material/input'; 
import {MatButtonModule} from '@angular/material/button'; 
import {MatCardModule} from '@angular/material/card'; 
import {MatToolbarModule} from '@angular/material/toolbar';
import { RegisterComponent } from './register/register.component'; 
import { AuthService } from './auth.service';
import { UploadComponent } from './upload/upload.component';
import { ProfileComponent } from './profile/profile.component';
import {MatSelectModule} from '@angular/material/select'; 
import { AuthGuard } from './auth.guard';
import { MatIconModule } from '@angular/material/icon';
import { HomeComponent } from './home/home.component';
import { RecentsComponent } from './recents/recents.component';
import { AccountComponent } from './account/account.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DisplayComponent } from './display/display.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    ProfileComponent,
    UploadComponent,
    HomeComponent,
    RecentsComponent,
    AccountComponent,
    DisplayComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule, 
    FlexLayoutModule, 
    MatFormFieldModule, 
    MatInputModule, 
    MatSelectModule,
    MatButtonModule, 
    MatCardModule, 
    HttpClientModule,
    MatToolbarModule,
    FormsModule,
    MatIconModule
  ],
  //schemas: [ NO_ERRORS_SCHEMA ],
  providers: [AuthService, AuthGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
