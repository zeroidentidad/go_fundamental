import { Component, OnInit } from '@angular/core';
import {MatDialogRef} from "@angular/material";
import {Employee} from "../../models/employee/employee";
import {Store} from '@ngrx/store';
import * as employeeActions from '../../state/actions/employee.actions';
import * as fromReducer from '../../state/reducers';

@Component({
  selector: 'app-employee-new-container',
  templateUrl: './employee-new-container.component.html',
  styleUrls: ['./employee-new-container.component.scss']
})
export class EmployeeNewContainerComponent implements OnInit {

  constructor(private dialogRef: MatDialogRef<EmployeeNewContainerComponent>, private store: Store<fromReducer.EmployeeState>) { }

  ngOnInit() {
  }

  onAdd(event: Employee): void {
    this.store.dispatch(new employeeActions.AddEmployee(event));
    this.dialogRef.close();
  }  

}
