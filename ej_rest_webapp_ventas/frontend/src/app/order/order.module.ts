import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { OrderRoutingModule } from './order-routing.module';
import { OrderMainContainerComponent } from './containers/order-main-container/order-main-container.component';
import {SharedModule} from "../shared/shared.module";
import {StoreModule} from "@ngrx/store";
import {EffectsModule} from "@ngrx/effects";
import {reducers} from './state/reducers';
import {OrderEffects} from "./state/effects/order.effects";


@NgModule({
  declarations: [OrderMainContainerComponent],
  imports: [
    CommonModule,
    OrderRoutingModule,
    SharedModule,
    StoreModule.forFeature('order', reducers),
    EffectsModule.forFeature([OrderEffects])
  ]
})
export class OrderModule { }
