import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {MatDialog} from "@angular/material";
import {CustomerPopupContainerComponent} from "../customer-popup-container/customer-popup-container.component";
import {Customer} from "../../models/customer/customer";
import {ProductPopupContainerComponent} from "../product-popup-container/product-popup-container.component";
import {Product} from "../../models/product/product";
import {PreOrderProduct} from "../../models/pre-order/pre-order-product";
import {PreOrderFooter} from "../../models/pre-order/pre-order-footer";

@Component({
  selector: 'app-order-manage-container',
  templateUrl: './order-manage-container.component.html',
  styleUrls: ['./order-manage-container.component.scss']
})
export class OrderManageContainerComponent implements OnInit {

  orderForm: FormGroup;
  orderProductList: PreOrderProduct[]=[];
  preOrderFooter: PreOrderFooter=PreOrderFooter.createEmptyInstance();
  constructor(private fb: FormBuilder, private dialog: MatDialog) { 
    this.buildNewForm();
  }

  ngOnInit() {
  }

  buildNewForm() {
    this.orderForm=this.fb.group({
      id: ['', [Validators.required]],
      name: ['', [Validators.required]],
      company: ['', [Validators.required]],
      fecha: [new Date(), [Validators.required]],
      address: ['', [Validators.required]],
      phone: ['', [Validators.required]],
      city: ['', [Validators.required]]
    });
  }

  openCustomerPopup(){
    const dialogRef = this.dialog.open(CustomerPopupContainerComponent, {
      panelClass: ''
    })
    
    dialogRef.afterClosed().subscribe((response: Customer)=>{
      if (response) {
        this.orderForm.patchValue({
          id: response.id,
          name: response.name,
          company: response.company,
          address: response.address,
          phone: response.businessphone,
          city: response.city,
        });        
      }
    })
  }

  openProductPopup(){
    const dialogRef=this.dialog.open(ProductPopupContainerComponent, {
      panelClass: ''
    });
    dialogRef.afterClosed()
      .subscribe((response: Product) => {
        this.AddProductToList(response);
      });
  }

  AddProductToList(item: Product): void {
    const product=new PreOrderProduct(0, item.id, item.description, item.productName, item.standardCost, 1);
    this.orderProductList.push(product);
    this.orderProductList=[...this.orderProductList];

    this.preOrderFooter=new PreOrderFooter(this.orderProductList);
  }  

  UpdateQuantity(event: any): any {
    const productToUpdate={...this.orderProductList[event.index]};
    productToUpdate.Quantity=Number(event.newValue);
    productToUpdate.TotalValue=Number((productToUpdate.Quantity*productToUpdate.UnitPrice).toFixed(2));

    this.orderProductList[event.index]=productToUpdate;
    this.orderProductList=[...this.orderProductList];

    this.preOrderFooter=new PreOrderFooter(this.orderProductList);
  }

  onDeleteProductOrder(orderDetailId: any) {

  }  

}
