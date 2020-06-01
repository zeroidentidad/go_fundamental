import { Component, OnInit } from '@angular/core';
import * as fromReducer from '../../state/reducers';
import * as customerActions from '../../state/actions/customer.actions';
import {Store} from '@ngrx/store';
import {GetCustomers} from '../../models/customer/get-customers';
import {Observable} from 'rxjs';
import {Customer} from "../../models/customer/customer";

@Component({
  selector: 'app-customer-popup-container',
  templateUrl: './customer-popup-container.component.html',
  styleUrls: ['./customer-popup-container.component.scss']
})
export class CustomerPopupContainerComponent implements OnInit {

  request: GetCustomers;
  pageSize=5;
  pageSizeOptions: number[]=[5, 10, 20, 30];  

  customers$: Observable<Customer[]>=this.store.select(fromReducer.getcustomers);
  length$: Observable<number>=this.store.select(fromReducer.gettotalCustomerRecords);
  
  constructor(private store: Store<fromReducer.OrderState>) { }

  ngOnInit() {
    this.request=new GetCustomers(this.pageSizeOptions[0], 0);
    this.store.dispatch(new customerActions.LoadCustomers(this.request));
  }

  changePage(event: any): void {
    const offSet=event.pageIndex*event.pageSize;
    this.pageSize=event.pageSize;
    this.request=new GetCustomers(event.pageSize, offSet);
    this.store.dispatch(new customerActions.LoadCustomers(this.request));
  }  

}
