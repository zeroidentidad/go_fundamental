import {Action} from "@ngrx/store";
import {GetEmployee} from "../../models/employee/get-employee";
import {EmployeeList} from "../../models/employee/employee-list";
import {Employee} from "../../models/employee/employee";
import {BestEmployee} from "../../models/employee/best-employee";

export enum EmployeeActionTypes {
    LoadEmployees = '[Employee] Load Employees',
    LoadEmployeesCompleted = '[Employee] Load Employees Completed',
    LoadEmployeeById = '[Employee] Load Employee By Id',
    LoadEmployeeByIdCompleted = '[Employee] Load Employee By Id Completed',
    AddEmployee = '[Employee] Add Employee',
    AddEmployeeCompleted = '[Employee] Add Employee Completed',
    DeleteEmployee = '[Employee] Delete Employee',
    DeleteEmployeeCompleted = '[Employee] Delete Employee Completed',
    LoadBestEmployee = '[Employee] Load Best Employee',
    LoadBestEmployeeCompleted = '[Employee] Load Best Employee Completed'
}

export class LoadEmployees implements Action {
    readonly type=EmployeeActionTypes.LoadEmployees;
    constructor(public request: GetEmployee) {}
}

export class LoadEmployeesCompleted implements Action {
    readonly type=EmployeeActionTypes.LoadEmployeesCompleted;
    constructor(public payload: EmployeeList) {}
}

export class LoadEmployeeById implements Action {
    readonly type=EmployeeActionTypes.LoadEmployeeById;
    constructor(public employeeId: number) {}
}

export class LoadEmployeeByIdCompleted implements Action {
    readonly type=EmployeeActionTypes.LoadEmployeeByIdCompleted;
    constructor(public payload: Employee) {}
}

export class AddEmployee implements Action {
    readonly type=EmployeeActionTypes.AddEmployee;
    constructor(public employee: Employee) {}
}

export class AddEmployeeCompleted implements Action {
    readonly type=EmployeeActionTypes.AddEmployeeCompleted;
    constructor() {}
}

export class DeleteEmployee implements Action {
    readonly type=EmployeeActionTypes.DeleteEmployee;
    constructor(public employeeId: number) {}
}
export class DeleteEmployeeCompleted implements Action {
    readonly type=EmployeeActionTypes.DeleteEmployeeCompleted;
    constructor() {}
}

export class LoadBestEmployee implements Action {
    readonly type=EmployeeActionTypes.LoadBestEmployee;
    constructor() {}
}

export class LoadBestEmployeeCompleted implements Action {
    readonly type=EmployeeActionTypes.LoadBestEmployeeCompleted;
    constructor(public payload: BestEmployee) {}
}

export type Actions=LoadEmployees|LoadEmployeesCompleted|LoadEmployeeById|LoadEmployeeByIdCompleted|AddEmployee|AddEmployeeCompleted|DeleteEmployee|DeleteEmployeeCompleted|LoadBestEmployee|LoadBestEmployeeCompleted;