import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {OrderMainContainerComponent} from "./containers/order-main-container/order-main-container.component";
import {OrderManageContainerComponent} from "./containers/order-manage-container/order-manage-container.component";


const routes: Routes = [
  {
    path: '',
    component: OrderMainContainerComponent
  },
  {
    path: 'manage',
    component: OrderManageContainerComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class OrderRoutingModule { }
