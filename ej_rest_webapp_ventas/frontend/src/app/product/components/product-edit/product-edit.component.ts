import { Component, OnInit, Input, SimpleChanges } from '@angular/core';
import {Product} from "../../models/product/product";
import {FormBuilder, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-product-edit',
  templateUrl: './product-edit.component.html',
  styleUrls: ['./product-edit.component.scss']
})
export class ProductEditComponent implements OnInit {

  @Input()
  product: Product;

  productEditForm: FormGroup;
  constructor(private fb: FormBuilder) { 

  }

  ngOnChanges(changes: SimpleChanges):void {
    if(changes.product&&changes.product.currentValue){
      this.buildProductEditForm()
    }
  }

  ngOnInit() {
  }

  buildProductEditForm():void {
    this.productEditForm = this.fb.group({
      productCode: [this.product.productCode, [Validators.required]],
      productName: [this.product.productName, [Validators.required]],
      category: [this.product.category, [Validators.required]],
      description: [this.product.description, [Validators.required]],
      standardCost: [this.product.standardCost, [Validators.required]],
      listPrice: [this.product.listPrice, [Validators.required]]
    });
  }

  onEdit():void{

  }

}
