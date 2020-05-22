import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatButtonModule, MatTabsModule, MatCardModule, MatPaginatorModule, MatIconModule, MatDialogModule} from '@angular/material';



@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    MatButtonModule,
    MatTabsModule,
    MatCardModule,
    MatPaginatorModule,
    MatIconModule, 
    MatDialogModule
  ],
  exports: [
    MatButtonModule,
    MatTabsModule,
    MatCardModule,
    MatPaginatorModule,
    MatIconModule,
    MatDialogModule
  ]
})
export class MaterialModule { }
