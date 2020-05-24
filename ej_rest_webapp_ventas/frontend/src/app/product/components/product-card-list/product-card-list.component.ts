import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import {Product} from "../../models/product/product";

@Component({
  selector: 'app-product-card-list',
  templateUrl: './product-card-list.component.html',
  styleUrls: ['./product-card-list.component.scss']
})
export class ProductCardListComponent implements OnInit {

  @Input()
  items: Product[];

  @Output()
  edit: EventEmitter<number>=new EventEmitter<number>();

  @Output()
  delete: EventEmitter<number>=new EventEmitter<number>();  

  constructor() { }

  ngOnInit() {
  }

  onEdit(productId: number): void {
    this.edit.emit(productId)
  }

  onDelete(productId: number): void {
    this.delete.emit(productId)
  }  

}
