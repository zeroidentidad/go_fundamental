import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProductRoutingModule } from './product-routing.module';
import { ProductMainContainerComponent } from './containers/product-main-container/product-main-container.component';


@NgModule({
  declarations: [ProductMainContainerComponent],
  imports: [
    CommonModule,
    ProductRoutingModule
  ]
})
export class ProductModule { }
