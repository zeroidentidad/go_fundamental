import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {GetOrder} from "../models/order/get-order";
import {Observable} from "rxjs";
import {OrderList} from "../models/order/order-list";
import {environment} from "src/environments/environment";
import {map} from "rxjs/operators";
import {OrderListItem} from "../models/order/order-list-item";

@Injectable({
  providedIn: 'root'
})
export class OrderService {

  constructor(private httpClient: HttpClient) { }

  getOrders(request: GetOrder): Observable<OrderList> {
    return this.httpClient.post<OrderList>(`${environment.ApiURL}orders/paginated`, request)
      .pipe(
        map((response: any) => {
          const orderList: OrderList = new OrderList();
          if (!response.data) {
            orderList.data = [];
            return orderList;
          }

          const orderListItem: OrderListItem[] = [];
          for (let index = 0; index < response.data.length; index++) {
            const element = response.data[index];
            orderListItem.push(OrderListItem.mapFromResponse(element))
          }

          orderList.data = orderListItem;
          orderList.totalRecords = response.totalRecords;
          return orderList;
        })
      );
  }

}
