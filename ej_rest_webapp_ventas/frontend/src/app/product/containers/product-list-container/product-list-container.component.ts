import { Component, OnInit } from '@angular/core';
import * as productActions from '../../state/actions/product.actions';
import * as fromReducer from '../../state/reducers'
import {GetProduct} from "../../models/product/get-product";
import {Store} from "@ngrx/store";
import {Product} from "../../models/product/product";
import {Observable} from "rxjs";

@Component({
  selector: 'app-product-list-container',
  templateUrl: './product-list-container.component.html',
  styleUrls: ['./product-list-container.component.scss']
})
export class ProductListContainerComponent implements OnInit {

  constructor(private store: Store<fromReducer.ProductState>) { }
  product$: Observable<Product[]>=this.store.select(fromReducer.getProducts)
  totalRecords$: Observable<number>=this.store.select(fromReducer.getTotalRecords)

  ngOnInit() {
    let request = new GetProduct(10,0);
    this.store.dispatch(new productActions.LoadProducts(request));
  }

}
