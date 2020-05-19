import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ProductRoutingModule } from './product-routing.module';
import { ProductMainContainerComponent } from './containers/product-main-container/product-main-container.component';
import { ProductListContainerComponent } from './containers/product-list-container/product-list-container.component';
import {SharedModule} from "../shared/shared.module";
import {StoreModule} from '@ngrx/store';
import {EffectsModule} from '@ngrx/effects';
import {reducers} from "../state/reducers";
import {ProductEffects} from "./state/effects/product.effects";


@NgModule({
  declarations: [ProductMainContainerComponent, ProductListContainerComponent],
  imports: [
    CommonModule,
    ProductRoutingModule,
    SharedModule,
    StoreModule.forFeature('product', reducers),
    EffectsModule.forFeature([ProductEffects])
  ]
})
export class ProductModule { }
