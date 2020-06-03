import { Component, OnInit, Input, ChangeDetectionStrategy } from '@angular/core';
import {Product} from "../../models/product/product";
import {MatDialogRef} from "@angular/material";

@Component({
  selector: 'app-product-popup',
  templateUrl: './product-popup.component.html',
  styleUrls: ['./product-popup.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ProductPopupComponent implements OnInit {

  @Input()
  products: Product;
  selectedProduct: Product;
  constructor(private dialogRef: MatDialogRef<ProductPopupComponent>) { }

  ngOnInit() {
  }

  onSelect(event: any, row: Product) {
    if(event.source.checked) {
      this.selectedProduct=row;
    }
  }  

  selectRow(): void {
    this.dialogRef.close(this.selectedProduct);
  }

}
