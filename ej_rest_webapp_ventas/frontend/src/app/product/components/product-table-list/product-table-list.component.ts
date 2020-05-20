import { Component, OnInit, Input } from '@angular/core';
import {Product} from "../../models/product/product";

@Component({
  selector: 'app-product-table-list',
  templateUrl: './product-table-list.component.html',
  styleUrls: ['./product-table-list.component.scss']
})
export class ProductTableListComponent implements OnInit {

  @Input()
  items: Product[]
  
  constructor() { }

  ngOnInit() {
  }

}
