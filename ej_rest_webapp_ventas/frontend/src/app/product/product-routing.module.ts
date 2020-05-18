import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {ProductMainContainerComponent} from "./containers/product-main-container/product-main-container.component";
import {ProductListContainerComponent} from "./containers/product-list-container/product-list-container.component";


const routes: Routes = [
  {
    path: '',
    component: ProductMainContainerComponent,
    children: [
      {
        path: '',
        redirectTo: 'list',
        pathMatch: 'full'
      },
      {
        path: 'list',
        component: ProductListContainerComponent,
      }
    ]
  }  
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProductRoutingModule { }
