import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ProjectCreateComponent } from './dialogs/project/project-create/project-create.component';
import { SatelliteCreateComponent } from './dialogs/satellite/satellite-create/satellite-create.component';
import { GroundStationCreateComponent } from './dialogs/groundStation/ground-station-create/ground-station-create.component';
import {HttpClientModule} from "@angular/common/http";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {ReactiveFormsModule} from "@angular/forms";
import {ButtonModule} from "primeng/button";
import {RadioButtonModule} from "primeng/radiobutton";

@NgModule({
  declarations: [
    AppComponent,
    ProjectCreateComponent,
    SatelliteCreateComponent,
    GroundStationCreateComponent
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
    ButtonModule,
    RadioButtonModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
