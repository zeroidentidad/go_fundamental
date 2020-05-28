import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { OrderRoutingModule } from './order-routing.module';
import { OrderMainContainerComponent } from './containers/order-main-container/order-main-container.component';
import {SharedModule} from "../shared/shared.module";


@NgModule({
  declarations: [OrderMainContainerComponent],
  imports: [
    CommonModule,
    OrderRoutingModule,
    SharedModule
  ]
})
export class OrderModule { }
