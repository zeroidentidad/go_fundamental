import { Component, OnInit, Input } from '@angular/core';
import {Customer} from "../../models/customer/customer";
import {MatDialogRef} from '@angular/material';

@Component({
  selector: 'app-customer-popup',
  templateUrl: './customer-popup.component.html',
  styleUrls: ['./customer-popup.component.scss']
})
export class CustomerPopupComponent implements OnInit {

  @Input()
  customers: Customer[];

  selectedCustomer: Customer;
  constructor(private dialogRef: MatDialogRef<CustomerPopupComponent>) { }

  ngOnInit() {
  }

  onSelect(event: any, row: Customer) {
    if(event.source.checked) {
      this.selectedCustomer=row;
    }
  } 
  
  selectRow(): void {
    this.dialogRef.close(this.selectedCustomer)
  }

}
