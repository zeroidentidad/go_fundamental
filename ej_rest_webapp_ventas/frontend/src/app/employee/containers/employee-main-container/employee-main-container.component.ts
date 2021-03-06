import { Component, OnInit } from '@angular/core';
import * as employeeActions from '../../state/actions/employee.actions';
import * as fromReducer from '../../state/reducers';
import {Store, ActionsSubject} from '@ngrx/store';
import {GetEmployee} from "../../models/employee/get-employee";
import {Employee} from "../../models/employee/employee";
import {Observable, Subject} from "rxjs";
import {MatDialog} from "@angular/material";
import {EmployeeDetailContainerComponent} from "../employee-detail-container/employee-detail-container.component";
import {EmployeeNewContainerComponent} from "../employee-new-container/employee-new-container.component";
import {ConfirmData} from "src/app/shared/models/confirm-data";
import {AppConfirmService} from "src/app/shared/components/app-confirm/app-confirm.service";
import {takeUntil} from "rxjs/operators";
import {ofType} from "@ngrx/effects";
import {BestEmployee} from "../../models/employee/best-employee";
import {EmployeeEditContainerComponent} from "../employee-edit-container/employee-edit-container.component";

@Component({
  selector: 'app-employee-main-container',
  templateUrl: './employee-main-container.component.html',
  styleUrls: ['./employee-main-container.component.scss']
})
export class EmployeeMainContainerComponent implements OnInit {

  request: GetEmployee;
  private ngUnsubscribe: Subject<any>=new Subject<any>();
  constructor(private store: Store<fromReducer.EmployeeState>, private dialog: MatDialog, private confirm: AppConfirmService, private actionSubject$: ActionsSubject) { 
    this.triggers();
   }

  employees$: Observable<Employee[]> = this.store.select(fromReducer.getEmployees);
  bestEmployee$: Observable<BestEmployee> = this.store.select(fromReducer.getBestEmployee);
  length$: Observable<number> = this.store.select(fromReducer.getTotalRecords);
  pageSize=5;
  pageSizeOptions: number[]=[5, 10, 25, 50];
 
  ngOnInit() {
    this.refreshData();
  }

  triggers(): void {
    this.actionSubject$
      .pipe(takeUntil(this.ngUnsubscribe))
      .pipe(ofType(employeeActions.EmployeeActionTypes.DeleteEmployeeCompleted))
      .subscribe(_ => {
        this.refreshData();
      })
  }

  refreshData(): void {
    this.request=new GetEmployee(this.pageSizeOptions[0], 0);
    this.store.dispatch(new employeeActions.LoadEmployees(this.request));
    this.store.dispatch(new employeeActions.LoadBestEmployee());
  } 
  
  changePage(event: any): void {
    const offSet=event.pageIndex*event.pageSize;
    this.request=new GetEmployee(event.pageSize, offSet);
    this.store.dispatch(new employeeActions.LoadEmployees(this.request));
  }  

  viewDetails(event: any): void {
    const dialogRef=this.dialog.open(EmployeeDetailContainerComponent,
      {
        panelClass: 'employee-modal-dialog',
        data: {employeeId: event.id}
      }
    );

    dialogRef.afterClosed().subscribe(_ => this.refreshData());
  }

  onAdd(): void {
    const dialogRef=this.dialog.open(EmployeeNewContainerComponent,
      {
        panelClass: 'employee-modal-dialog'
      }
    );

    dialogRef.afterClosed().subscribe(_ => this.refreshData());
  }  

  onDelete(event: any): void {
    const confirmData=new ConfirmData('Eliminar empleado', 'Estas seguro de eliminar el empleado?', true);
    this.confirm.confirm(confirmData)
      .subscribe(result => {
        if(result) {
          this.store.dispatch(new employeeActions.DeleteEmployee(event.id));
        }
      })
  } 

  onEdit(employeeId: number): void {
    const dialogRef=this.dialog.open(EmployeeEditContainerComponent, {
      panelClass: 'employee-modal-dialog',
      data: {employeeId: employeeId}
    });

    dialogRef.afterClosed().subscribe(result => {
      this.refreshData();
    });
  }   

}
