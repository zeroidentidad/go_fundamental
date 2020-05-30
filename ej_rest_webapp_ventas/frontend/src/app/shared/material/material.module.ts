import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {MatButtonModule, MatTabsModule, MatCardModule, MatPaginatorModule, MatIconModule, MatDialogModule, MatFormFieldModule, MatInputModule, MatSelectModule, MatDividerModule, MatTooltipModule, MatExpansionModule, MatListModule, MatNativeDateModule, MatDatepickerModule, MatRadioModule} from '@angular/material';



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
    MatTooltipModule,
    MatExpansionModule,
    MatListModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatRadioModule
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
    MatTooltipModule,
    MatExpansionModule,
    MatListModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatRadioModule
  ]
})
export class MaterialModule { }
