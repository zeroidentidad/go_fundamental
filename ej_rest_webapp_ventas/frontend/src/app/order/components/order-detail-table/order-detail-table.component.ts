import { Component, OnInit, Input } from '@angular/core';
import {PreOrderProduct} from "../../models/pre-order/pre-order-product";

@Component({
  selector: 'app-order-detail-table',
  templateUrl: './order-detail-table.component.html',
  styleUrls: ['./order-detail-table.component.scss']
})
export class OrderDetailTableComponent implements OnInit {

  @Input()
  items: PreOrderProduct[];
  constructor() { }

  ngOnInit() {
  }

}
