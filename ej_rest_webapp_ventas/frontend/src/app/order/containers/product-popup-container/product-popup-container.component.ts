import { Component, OnInit } from '@angular/core';
import * as productActions from '../../state/actions/product.actions';
import * as fromReducer from '../../state/reducers';
import {Store} from '@ngrx/store';
import {GetProducts} from "../../models/product/get-product";
import {Observable} from "rxjs";
import {Product} from "../../models/product/product";

@Component({
  selector: 'app-product-popup-container',
  templateUrl: './product-popup-container.component.html',
  styleUrls: ['./product-popup-container.component.scss']
})
export class ProductPopupContainerComponent implements OnInit {

  products$: Observable<Product[]>=this.store.select(fromReducer.getProducts);
  length$: Observable<number>=this.store.select(fromReducer.gettotalProductRecords);

  pageSize=5;
  pageSizeOptions: number[]=[5, 10, 20, 30];
  request: GetProducts;
  constructor(private store: Store<fromReducer.OrderState>) { }

  ngOnInit() {
    this.request=new GetProducts(this.pageSizeOptions[0], 0);
    this.store.dispatch(new productActions.LoadProducts(this.request));
  }

  changePage(event: any): void {
    const offSet=event.pageIndex*event.pageSize;
    this.pageSize=event.pageSize;
    this.request=new GetProducts(event.pageSize, offSet);
    this.store.dispatch(new productActions.LoadProducts(this.request));
  }  

}
