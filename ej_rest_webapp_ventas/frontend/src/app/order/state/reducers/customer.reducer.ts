import { Customer } from '../../models/customer/customer';
import * as customerActions from '../actions/customer.actions';

export interface State {
    customers: Customer[];
    totalCustomerRecords: number;
}

const initialState: State = {
    customers: [],
    totalCustomerRecords: 0
};

export function CustomerReducer(state = initialState, action: customerActions.Actions): State {
    switch (action.type) {
        case customerActions.CustomerActionTypes.LoadCustomers: {
            return state;
        }
        case customerActions.CustomerActionTypes.LoadCustomersCompleted: {
            return {
                ...state,
                customers: action.payload.data,
                totalCustomerRecords: action.payload.totalRecords
            };
        }
        default: return state;

    }

}
export const getCustomers = (state: State) => state.customers;
export const getTotalCustomerRecords = (state: State) => state.totalCustomerRecords;