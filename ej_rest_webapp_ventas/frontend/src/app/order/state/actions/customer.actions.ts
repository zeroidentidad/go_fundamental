import {Action} from '@ngrx/store';
import {GetCustomers} from '../../models/customer/gest-customers';
import {CustomerList} from '../../models/customer/customer-list';

export enum CustomerActionTypes {
    LoadCustomers='[Order-Customer] Load Customers',
    LoadCustomersCompleted='[Order-Customer] Load Customers Completed'
}

export class LoadCustomers implements Action {
    readonly type=CustomerActionTypes.LoadCustomers;
    constructor(public request: GetCustomers) {}
}

export class LoadCustomersCompleted implements Action {
    readonly type=CustomerActionTypes.LoadCustomersCompleted;
    constructor(public payload: CustomerList) {}
}

export type Actions=LoadCustomers|LoadCustomersCompleted;