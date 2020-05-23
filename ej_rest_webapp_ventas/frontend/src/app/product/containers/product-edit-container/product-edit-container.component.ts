import { Component, OnInit, Inject } from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material";
import {Store} from "@ngrx/store";
import * as fromReducer from "../../state/reducers";
import * as productActions from "../../state/actions/product.actions";
import {Product} from "../../models/product/product";
import {Observable} from "rxjs";

interface ProductEditData {
  productId: number;
}


@Component({
  selector: 'app-product-edit-container',
  templateUrl: './product-edit-container.component.html',
  styleUrls: ['./product-edit-container.component.scss']
})
export class ProductEditContainerComponent implements OnInit {

  product$: Observable<Product> = this.store.select(fromReducer.getProduct);
  constructor(@Inject(MAT_DIALOG_DATA) private data: ProductEditData,
    private dialogRef: MatDialogRef<ProductEditContainerComponent>,
    private store: Store<fromReducer.ProductState>) { 
    this.store.dispatch(new productActions.GetProductById(data.productId))
     }

  ngOnInit() {
  }

  onEdit(product: Product): void {
    this.store.dispatch(new productActions.UpdateProduct(product))
    this.dialogRef.close()
  }

}
