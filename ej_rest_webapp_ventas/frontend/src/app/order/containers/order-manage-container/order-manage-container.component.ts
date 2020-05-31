import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";

@Component({
  selector: 'app-order-manage-container',
  templateUrl: './order-manage-container.component.html',
  styleUrls: ['./order-manage-container.component.scss']
})
export class OrderManageContainerComponent implements OnInit {

  orderForm: FormGroup;
  constructor(private fb: FormBuilder) { 
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
    
  }

}
