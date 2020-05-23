import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-product-new-container',
  templateUrl: './product-new-container.component.html',
  styleUrls: ['./product-new-container.component.scss']
})
export class ProductNewContainerComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

  buildCategories(): any {
    return [
      {key: 'Baked Goods & Mixes', value: 'Baked Goods & Mixes'},
      {key: 'Beverages', value: 'Beverages'},
      {key: 'Candy', value: 'Candy'},
      {key: 'Canned Fruit & Vegetables', value: 'Canned Fruit & Vegetables'},
      {key: 'Canned Meat', value: 'Canned Meat'},
      {key: 'Cereal', value: 'Cereal'},
      {key: 'Chips, Snacks', value: 'Chips, Snacks'},
      {key: 'Condiments', value: 'Condiments'},
      {key: 'Grains', value: 'Grains'},
      {key: 'Oil', value: 'Oil'},
      {key: 'Pasta', value: 'Pasta'},
      {key: 'Sauces', value: 'Sauces'},
      {key: 'Soups', value: 'Soups'}
    ];
  }  

}
