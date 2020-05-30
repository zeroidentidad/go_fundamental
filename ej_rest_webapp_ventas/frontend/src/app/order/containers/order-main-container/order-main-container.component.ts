import { Component, OnInit, TemplateRef, ViewChild, AfterViewInit } from '@angular/core';
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
export class OrderMainContainerComponent implements OnInit, AfterViewInit {

  orders$: Observable<OrderListItem[]>=this.store.select(fromReducer.getOrders);
  length$: Observable<number>=this.store.select(fromReducer.getTotalRecords);  

  request: GetOrder;
  pageSize=10;
  pageSizeOptions: number[]=[10, 20, 50, 100];

  columns: object[]=[];

  @ViewChild("orderDateCellTemplate", {static: false})
  private orderDateCellTemplate: TemplateRef<any>;

  constructor(private store: Store<fromReducer.OrderState>,) { 
    this.refreshData();
  }

  ngOnInit() {
  }

  ngAfterViewInit(): void {
    this.columns=this.getColumns();
  }

  refreshData(): void {
    this.request=new GetOrder(this.pageSizeOptions[0], 0);
    this.store.dispatch(new orderActions.LoadOrders(this.request));
  }

  getColumns(): object[] {
    return [
      {
        name: "Id",
        prop: "orderId"
      },
      {
        name: "Cliente",
        prop: "customer"
      },
      {
        name: "Telefono",
        prop: "phone"
      },
      {
        name: "Dirección",
        prop: "address"
      },
      {
        name: "Ciudad",
        prop: "city"
      },
      {
        name: "Fecha",
        cellTemplate: this.orderDateCellTemplate
      }
    ];
  }  

}
