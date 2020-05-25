import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {EmployeeMainContainerComponent} from "./containers/employee-main-container/employee-main-container.component";


const routes: Routes = [
  {
    path:'',
    component:EmployeeMainContainerComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class EmployeeRoutingModule { }

