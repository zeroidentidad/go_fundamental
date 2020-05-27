import * as employeeActions from '../actions/employee.actions';
import {Injectable} from '@angular/core';
import {EmployeeService} from '../../services/employee.service';
import {Actions, Effect, ofType} from '@ngrx/effects';
import {switchMap, map} from 'rxjs/operators';

@Injectable()
export class EmployeeEffects {
    constructor(private employeeService: EmployeeService,
        private actions$: Actions) {}

    @Effect()
    getEmployees$=this.actions$.pipe(
        ofType<employeeActions.LoadEmployees>(employeeActions.EmployeeActionTypes.LoadEmployees),
        switchMap(action => this.employeeService.getEmployees(action.request)
            .pipe(
                map(data => new employeeActions.LoadEmployeesCompleted(data))
            ))
    );

    @Effect()
    getEmployeesById$=this.actions$.pipe(
        ofType<employeeActions.LoadEmployeeById>(employeeActions.EmployeeActionTypes.LoadEmployeeById),
        switchMap(action => this.employeeService.getEmployeeById(action.employeeId)
            .pipe(
                map(data => new employeeActions.LoadEmployeeByIdCompleted(data))
            ))
    );   

    @Effect()
    addEmployee$=this.actions$.pipe(
        ofType<employeeActions.AddEmployee>(employeeActions.EmployeeActionTypes.AddEmployee),
        switchMap(action => this.employeeService.createEmployee(action.employee)
            .pipe(
                map(_ => new employeeActions.AddEmployeeCompleted())
            ))
    );
    
    @Effect()
    deleteEmployees$=this.actions$.pipe(
        ofType<employeeActions.DeleteEmployee>(employeeActions.EmployeeActionTypes.DeleteEmployee),
        switchMap(action => this.employeeService.deleteEmployee(action.employeeId)
            .pipe(
                map(_ => new employeeActions.DeleteEmployeeCompleted())
            ))
    ); 

}