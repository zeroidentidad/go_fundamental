import { Component, OnInit } from '@angular/core';
import * as productActions from '../../state/actions/product.actions';
import * as fromReducer from '../../state/reducers'
import {GetProduct} from "../../models/product/get-product";
import {Store, ActionsSubject} from "@ngrx/store";
import {Product} from "../../models/product/product";
import {Observable} from "rxjs";
import {MatDialog} from "@angular/material";
import {ProductEditContainerComponent} from "../product-edit-container/product-edit-container.component";
import {ConfirmData} from "src/app/shared/models/confirm-data";
import {AppConfirmService} from "src/app/shared/components/app-confirm/app-confirm.service";
import {ofType} from "@ngrx/effects";

@Component({
  selector: 'app-product-list-container',
  templateUrl: './product-list-container.component.html',
  styleUrls: ['./product-list-container.component.scss']
})
export class ProductListContainerComponent implements OnInit {

  constructor(private store: Store<fromReducer.ProductState>, public dialog: MatDialog, private confirm: AppConfirmService, private actionsSubject$: ActionsSubject) { }
  products$: Observable<Product[]>=this.store.select(fromReducer.getProducts)
  totalRecords$: Observable<number>=this.store.select(fromReducer.getTotalRecords)
  pageSize = 10;
  pageSizeOptions: number[] = [10, 20, 30, 50];
  request: GetProduct;

  ngOnInit() {
    this.refreshData();
  }

  refreshData():void{
    this.request=new GetProduct(this.pageSizeOptions[0],0);
    this.store.dispatch(new productActions.LoadProducts(this.request));
  }

  triggers(): void {
    this.actionsSubject$
      .pipe(ofType(productActions.ProductActionTypes.DeleteProductCompleted))
      .subscribe(_ => this.refreshData())
  }

  changePage(event: any): void {
    const offSet = event.pageIndex * event.pageSize;
    this.pageSize = event.pageSize;
    this.request=new GetProduct(this.pageSize, offSet);
    this.store.dispatch(new productActions.LoadProducts(this.request));
  }

  onEdit(productId: number): void {
    const dialogRef=this.dialog.open(ProductEditContainerComponent, {
      width: '250px',
      data: {productId: productId}
    });

    dialogRef.afterClosed().subscribe(result => {
      this.refreshData();
    });
  }

  onDelete(productId: number): void {
    const confirmData=new ConfirmData("Eliminar producto", "Estas seguro de Eliminar?", true);
    this.confirm.confirm(confirmData)
      .subscribe(result => {
        if(result) {
          this.store.dispatch(new productActions.DeleteProduct(productId));
        }
      });
  }  

}
