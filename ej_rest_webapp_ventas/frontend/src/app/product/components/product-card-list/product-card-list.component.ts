import { Component, OnInit, Input } from '@angular/core';
import {Product} from "../../models/product/product";

@Component({
  selector: 'app-product-card-list',
  templateUrl: './product-card-list.component.html',
  styleUrls: ['./product-card-list.component.scss']
})
export class ProductCardListComponent implements OnInit {

  @Input()
  items: Product[];

  constructor() { }

  ngOnInit() {
  }

}
