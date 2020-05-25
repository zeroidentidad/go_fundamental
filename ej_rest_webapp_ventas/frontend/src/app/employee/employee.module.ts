import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { EmployeeRoutingModule } from './employee-routing.module';
import { EmployeeMainContainerComponent } from './containers/employee-main-container/employee-main-container.component';
import {SharedModule} from "../shared/shared.module";


@NgModule({
  declarations: [EmployeeMainContainerComponent],
  imports: [
    CommonModule,
    EmployeeRoutingModule,
    SharedModule
  ]
})
export class EmployeeModule { }
