import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import {RouterModule} from "@angular/router";
import { MaterialModule } from './material/material.module';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import { AppConfirmComponent } from './components/app-confirm/app-confirm.component';
import {AppConfirmService} from "./components/app-confirm/app-confirm.service";



@NgModule({
  declarations: [NavBarComponent, AppConfirmComponent],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    MaterialModule
  ],
  exports: [NavBarComponent, MaterialModule, FormsModule, ReactiveFormsModule, AppConfirmComponent],
  entryComponents: [AppConfirmComponent],
  providers: [AppConfirmService]
})
export class SharedModule { }
