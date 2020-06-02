import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {MatDialog} from "@angular/material";
import {CustomerPopupContainerComponent} from "../customer-popup-container/customer-popup-container.component";
import {Customer} from "../../models/customer/customer";
import {ProductPopupContainerComponent} from "../product-popup-container/product-popup-container.component";
import {Product} from "../../models/product/product";

@Component({
  selector: 'app-order-manage-container',
  templateUrl: './order-manage-container.component.html',
  styleUrls: ['./order-manage-container.component.scss']
})
export class OrderManageContainerComponent implements OnInit {

  orderForm: FormGroup;
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
        console.log(response);
      });
  }

  AddProductToList(item: Product): void {
    console.log('prro :v')
  }  

}
