import * as employeeActions from '../actions/employee.actions';
import {Injectable} from '@angular/core';
import {EmployeeService} from '../../services/employee.service';
import {Actions, Effect, ofType} from '@ngrx/effects';
import {switchMap, map} from 'rxjs/operators';
import {Router} from "@angular/router";

@Injectable()
export class EmployeeEffects {
    constructor(private employeeService: EmployeeService, private actions$: Actions, private router: Router) {}

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

    @Effect()
    getBestEmployee$=this.actions$.pipe(
        ofType<employeeActions.LoadBestEmployee>(employeeActions.EmployeeActionTypes.LoadBestEmployee),
        switchMap(_ => this.employeeService.getBestEmployee()
            .pipe(
                map(data => new employeeActions.LoadBestEmployeeCompleted(data))
            ))
    );

    @Effect()
    updateEmployee$=this.actions$.pipe(
        ofType<employeeActions.UpdateEmployee>(employeeActions.EmployeeActionTypes.UpdateEmployee),
        switchMap(action => this.employeeService.updateEmployee(action.request)
            .pipe(
                map(_ => {
                    this.router.navigate(['/employee'])
                    return new employeeActions.UpdateEmployeeCompleted()
                }
                )))
    );    

}