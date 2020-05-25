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

}