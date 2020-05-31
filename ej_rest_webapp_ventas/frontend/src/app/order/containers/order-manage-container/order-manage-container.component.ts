import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {MatDialog} from "@angular/material";
import {CustomerPopupContainerComponent} from "../customer-popup-container/customer-popup-container.component";

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
    
    dialogRef.afterClosed().subscribe(_=>{
      console.log('cerrado');
    })
  }

}
