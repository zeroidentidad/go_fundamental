import * as employeeActions from '../actions/employee.actions';
import {Employee} from "../../models/employee/employee";


export interface State {
    employees: Employee[];
    totalRecords: number;
    employee: Employee;
}

const initialState: State={
    employees: [],
    totalRecords: 0,
    employee: new Employee(),
}

export function EmployeeReducer(state=initialState, action: employeeActions.Actions): State {
    switch(action.type) {
        case employeeActions.EmployeeActionTypes.LoadEmployees:
            {
                return state;
            }

        case employeeActions.EmployeeActionTypes.LoadEmployeesCompleted:
            return {
                ...state,
                employees: action.payload.data,
                totalRecords: action.payload.totalRecords
            }

        case employeeActions.EmployeeActionTypes.LoadEmployeeById:
            return state;

        case employeeActions.EmployeeActionTypes.LoadEmployeeByIdCompleted:
            return {
                ...state,
                employee: action.payload
            }            

        default:
            return state;
    }
}

export const getEmployees=(state: State) => state.employees;
export const getTotalRecords=(state: State) => state.totalRecords;
export const getEmployee=(state: State) => state.employee;