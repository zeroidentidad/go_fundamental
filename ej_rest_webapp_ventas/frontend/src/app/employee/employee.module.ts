import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { EmployeeRoutingModule } from './employee-routing.module';
import { EmployeeMainContainerComponent } from './containers/employee-main-container/employee-main-container.component';
import {SharedModule} from "../shared/shared.module";
import {StoreModule} from "@ngrx/store";
import {reducers} from "./state/reducers";
import {EffectsModule} from "@ngrx/effects";
import {EmployeeEffects} from "./state/effects/employee.effects";
import { EmployeeListComponent } from './components/employee-list/employee-list.component';
import { EmployeeDetailContainerComponent } from './containers/employee-detail-container/employee-detail-container.component';
import { EmployeeDetailComponent } from './components/employee-detail/employee-detail.component';
import { EmployeeNewContainerComponent } from './containers/employee-new-container/employee-new-container.component';
import { EmployeeNewComponent } from './components/employee-new/employee-new.component';
import { BestEmployeeComponent } from './components/best-employee/best-employee.component';
import { EmployeeEditContainerComponent } from './containers/employee-edit-container/employee-edit-container.component';
import { EmployeeEditComponent } from './components/employee-edit/employee-edit.component';


@NgModule({
  declarations: [EmployeeMainContainerComponent, EmployeeListComponent, EmployeeDetailContainerComponent, EmployeeDetailComponent, EmployeeNewContainerComponent, EmployeeNewComponent, BestEmployeeComponent, EmployeeEditContainerComponent, EmployeeEditComponent],
  imports: [
    CommonModule,
    EmployeeRoutingModule,
    SharedModule,
    StoreModule.forFeature('employee', reducers),
    EffectsModule.forFeature([EmployeeEffects])
  ],
  entryComponents: [EmployeeDetailContainerComponent, EmployeeNewContainerComponent, EmployeeEditContainerComponent]
})
export class EmployeeModule { }
