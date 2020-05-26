import { Component, OnInit, Inject } from '@angular/core';
import {MAT_DIALOG_DATA} from "@angular/material";
import {Store} from '@ngrx/store';
import * as employeeActions from '../../state/actions/employee.actions';
import * as fromReducer from '../../state/reducers';
import {Employee} from "../../models/employee/employee";
import {Observable} from "rxjs";

interface DialogData {
  employeeId: number;
}

@Component({
  selector: 'app-employee-detail-container',
  templateUrl: './employee-detail-container.component.html',
  styleUrls: ['./employee-detail-container.component.scss']
})
export class EmployeeDetailContainerComponent implements OnInit {

  employee$: Observable<Employee>=this.store.select(fromReducer.getEmployee);

  constructor(@Inject(MAT_DIALOG_DATA) public data: DialogData, private store: Store<fromReducer.EmployeeState>) {
    this.store.dispatch(new employeeActions.LoadEmployeeById(data.employeeId));
  }

  ngOnInit() {
  }

}
