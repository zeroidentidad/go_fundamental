import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProductRoutingModule } from './product-routing.module';
import { ProductMainContainerComponent } from './containers/product-main-container/product-main-container.component';
import { ProductListContainerComponent } from './containers/product-list-container/product-list-container.component';
import {SharedModule} from "../shared/shared.module";
import {StoreModule} from '@ngrx/store';
import {EffectsModule} from '@ngrx/effects';
import {reducers} from "./state/reducers";
import {ProductEffects} from "./state/effects/product.effects";
import { ProductTableListComponent } from './components/product-table-list/product-table-list.component';
import { ProductCardListComponent } from './components/product-card-list/product-card-list.component';
import { ProductEditContainerComponent } from './containers/product-edit-container/product-edit-container.component';
import { ProductEditComponent } from './components/product-edit/product-edit.component';


@NgModule({
  declarations: [ProductMainContainerComponent, ProductListContainerComponent, ProductTableListComponent, ProductCardListComponent, ProductEditContainerComponent, ProductEditComponent],
  imports: [
    CommonModule,
    ProductRoutingModule,
    SharedModule,
    StoreModule.forFeature('product', reducers),
    EffectsModule.forFeature([ProductEffects])
  ],
  entryComponents: [ProductEditContainerComponent]
})
export class ProductModule { }
