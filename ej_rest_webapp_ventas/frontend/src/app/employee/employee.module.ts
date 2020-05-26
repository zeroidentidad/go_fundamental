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


@NgModule({
  declarations: [EmployeeMainContainerComponent, EmployeeListComponent, EmployeeDetailContainerComponent, EmployeeDetailComponent],
  imports: [
    CommonModule,
    EmployeeRoutingModule,
    SharedModule,
    StoreModule.forFeature('employee', reducers),
    EffectsModule.forFeature([EmployeeEffects])
  ],
  entryComponents: [EmployeeDetailContainerComponent]
})
export class EmployeeModule { }
