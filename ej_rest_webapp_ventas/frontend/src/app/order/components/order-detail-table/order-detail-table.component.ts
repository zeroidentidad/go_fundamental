import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import {PreOrderProduct} from "../../models/pre-order/pre-order-product";

@Component({
  selector: 'app-order-detail-table',
  templateUrl: './order-detail-table.component.html',
  styleUrls: ['./order-detail-table.component.scss']
})
export class OrderDetailTableComponent implements OnInit {

  @Input()
  items: PreOrderProduct[];
  @Output()
  updateQuantity: EventEmitter<any>=new EventEmitter<any>();
  @Output()
  deleteProductOrder: EventEmitter<any>=new EventEmitter<any>();

  editing={};
  constructor() { }

  ngOnInit() {
  }

  onRemoveItem(event: any): void {
    this.deleteProductOrder.emit(event);
  }  

  onUpdateQuantity(event: any, cell: any, rowIndex: any) {
    this.editing[rowIndex+'-'+cell]=false;
    const selectedObject={index: rowIndex, newValue: event.target.value};
    this.updateQuantity.emit(selectedObject);
  }  

}
