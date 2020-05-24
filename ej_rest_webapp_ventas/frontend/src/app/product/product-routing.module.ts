import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {ProductMainContainerComponent} from "./containers/product-main-container/product-main-container.component";
import {ProductListContainerComponent} from "./containers/product-list-container/product-list-container.component";
import {ProductNewContainerComponent} from "./containers/product-new-container/product-new-container.component";
import {ProductNewGuard} from "./guards/product-new.guard";


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
      },
      {
        path: 'new',
        component: ProductNewContainerComponent,
        canDeactivate: [ProductNewGuard]
      }
    ]
  }  
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ProductRoutingModule { }
