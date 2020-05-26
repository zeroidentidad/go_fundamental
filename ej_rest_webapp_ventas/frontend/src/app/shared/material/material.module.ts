import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatButtonModule, MatTabsModule, MatCardModule, MatPaginatorModule, MatIconModule, MatDialogModule, MatFormFieldModule, MatInputModule, MatSelectModule, MatDividerModule, MatTooltipModule} from '@angular/material';



@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    MatButtonModule,
    MatTabsModule,
    MatCardModule,
    MatPaginatorModule,
    MatIconModule, 
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDividerModule,
    MatTooltipModule
  ],
  exports: [
    MatButtonModule,
    MatTabsModule,
    MatCardModule,
    MatPaginatorModule,
    MatIconModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatDividerModule,
    MatTooltipModule
  ]
})
export class MaterialModule { }
