import { Component, OnInit, Inject } from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material";
import {Employee} from "../../models/employee/employee";
import {Observable} from "rxjs";
import * as fromReducer from "../../state/reducers";
import * as employeeActions from "../../state/actions/employee.actions";
import {Store} from "@ngrx/store";

interface EmployeeEditData {
  employeeId: number;
}

@Component({
  selector: 'app-employee-edit-container',
  templateUrl: './employee-edit-container.component.html',
  styleUrls: ['./employee-edit-container.component.scss']
})
export class EmployeeEditContainerComponent implements OnInit {

  employee$: Observable<Employee>=this.store.select(fromReducer.getEmployee);
  constructor(@Inject(MAT_DIALOG_DATA) private data: EmployeeEditData,
    private dialogRef: MatDialogRef<EmployeeEditContainerComponent>,
    private store: Store<fromReducer.EmployeeState>) {
    this.store.dispatch(new employeeActions.LoadEmployeeById(data.employeeId))
  }

  ngOnInit() {
  }

  onEdit(employee: Employee): void {
    this.store.dispatch(new employeeActions.UpdateEmployee(employee))
    this.dialogRef.close()
  }  

}
