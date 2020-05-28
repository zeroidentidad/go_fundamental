import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {OrderMainContainerComponent} from "./containers/order-main-container/order-main-container.component";


const routes: Routes = [
  {
    path: '',
    component: OrderMainContainerComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class OrderRoutingModule { }
