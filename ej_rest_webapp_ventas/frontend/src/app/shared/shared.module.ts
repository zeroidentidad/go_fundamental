import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import {RouterModule} from "@angular/router";
import { MaterialModule } from './material/material.module';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";



@NgModule({
  declarations: [NavBarComponent],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    MaterialModule
  ],
  exports: [NavBarComponent, MaterialModule, FormsModule, ReactiveFormsModule]
})
export class SharedModule { }
