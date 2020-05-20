import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatButtonModule, MatTabsModule, MatCardModule} from '@angular/material';



@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    MatButtonModule,
    MatTabsModule,
    MatCardModule
  ],
  exports: [
    MatButtonModule,
    MatTabsModule,
    MatCardModule
  ]
})
export class MaterialModule { }
