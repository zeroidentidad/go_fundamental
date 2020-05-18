import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import {RouterModule} from "@angular/router";
import { MaterialModule } from './material/material.module';



@NgModule({
  declarations: [NavBarComponent],
  imports: [
    CommonModule,
    RouterModule,
    MaterialModule
  ],
  exports: [NavBarComponent, MaterialModule]
})
export class SharedModule { }
