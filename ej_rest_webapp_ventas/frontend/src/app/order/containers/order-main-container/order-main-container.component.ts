import { Component, OnInit } from '@angular/core';
import {Store} from '@ngrx/store';
import * as orderActions from '../../state/actions/order.actions';
import * as fromReducer from '../../state/reducers';
import {GetOrder} from "../../models/order/get-order";
import {Observable} from "rxjs";
import {OrderListItem} from "../../models/order/order-list-item";

@Component({
  selector: 'app-order-main-container',
  templateUrl: './order-main-container.component.html',
  styleUrls: ['./order-main-container.component.scss']
})
export class OrderMainContainerComponent implements OnInit {

  orders$: Observable<OrderListItem[]>=this.store.select(fromReducer.getOrders);
  length$: Observable<number>=this.store.select(fromReducer.getTotalRecords);  

  request: GetOrder;
  pageSize=10;
  pageSizeOptions: number[]=[10, 20, 50, 100];

  constructor(private store: Store<fromReducer.OrderState>,) { 
    this.refreshData();
  }

  ngOnInit() {
  }

  refreshData(): void {
    this.request=new GetOrder(this.pageSizeOptions[0], 0);
    this.store.dispatch(new orderActions.LoadOrders(this.request));
  }


}
