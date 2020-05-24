import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import {RouterModule} from "@angular/router";
import { MaterialModule } from './material/material.module';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import { AppConfirmComponent } from './components/app-confirm/app-confirm.component';
import {AppConfirmService} from "./components/app-confirm/app-confirm.service";
import { DecimalOnlyDirective } from './directives/decimal-only.directive';
import {HighchartsChartModule} from 'highcharts-angular';

@NgModule({
  declarations: [NavBarComponent, AppConfirmComponent, DecimalOnlyDirective],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    MaterialModule,
    HighchartsChartModule
  ],
  exports: [NavBarComponent, MaterialModule, FormsModule, ReactiveFormsModule, AppConfirmComponent, DecimalOnlyDirective, HighchartsChartModule],
  entryComponents: [AppConfirmComponent],
  providers: [AppConfirmService]
})
export class SharedModule { }
