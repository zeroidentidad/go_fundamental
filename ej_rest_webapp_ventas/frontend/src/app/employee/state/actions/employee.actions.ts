import {Action} from "@ngrx/store";
import {GetEmployee} from "../../models/employee/get-employee";
import {EmployeeList} from "../../models/employee/employee-list";

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

export type Actions=LoadEmployees|LoadEmployeesCompleted;