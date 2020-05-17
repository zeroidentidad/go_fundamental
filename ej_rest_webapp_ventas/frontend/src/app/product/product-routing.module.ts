import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {ProductMainContainerComponent} from "./containers/product-main-container/product-main-container.component";


const routes: Routes = [
  {
    path: '',
    component: ProductMainContainerComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProductRoutingModule { }
