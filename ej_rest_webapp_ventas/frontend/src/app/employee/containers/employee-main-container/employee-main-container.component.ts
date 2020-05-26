import { Component, OnInit } from '@angular/core';
import * as employeeActions from '../../state/actions/employee.actions';
import * as fromReducer from '../../state/reducers';
import {Store, ActionsSubject} from '@ngrx/store';
import {GetEmployee} from "../../models/employee/get-employee";
import {Employee} from "../../models/employee/employee";
import {Observable} from "rxjs";

@Component({
  selector: 'app-employee-main-container',
  templateUrl: './employee-main-container.component.html',
  styleUrls: ['./employee-main-container.component.scss']
})
export class EmployeeMainContainerComponent implements OnInit {

  request: GetEmployee;
  constructor(private store:Store<fromReducer.EmployeeState>) { }

  employees$: Observable<Employee[]> = this.store.select(fromReducer.getEmployees);
  length$: Observable<number> = this.store.select(fromReducer.getTotalRecords);

  ngOnInit() {
    this.refreshdata();
  }

  pageSizeOptions: number[]=[5, 10, 25, 50];
  refreshdata(): void {
    this.request=new GetEmployee(this.pageSizeOptions[0], 0);
    this.store.dispatch(new employeeActions.LoadEmployees(this.request))
  }  

}
