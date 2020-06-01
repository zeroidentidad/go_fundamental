import {ActionReducerMap, createFeatureSelector, createSelector} from '@ngrx/store';
import * as fromOrderReducer from '../reducers/order.reducer';
import * as fromCustomerReducer from '../reducers/customer.reducer';

export interface OrderState {
    order: fromOrderReducer.State;
    customer: fromCustomerReducer.State;
}

export const reducers: ActionReducerMap<OrderState>={
    order: fromOrderReducer.OrderReducer,
    customer: fromCustomerReducer.CustomerReducer
}

export const getOrderModuleState=createFeatureSelector<OrderState>('order');

/// Order State
export const getOrderState=createSelector(getOrderModuleState, state => state.order);
export const getQuery=createSelector(getOrderState, fromOrderReducer.getQuery);
export const getOrders=createSelector(getOrderState, fromOrderReducer.getOrders);
export const getTotalRecords=createSelector(getOrderState, fromOrderReducer.getTotalOrderRecords);

/// Customer State
export const getCustomerState=createSelector(getOrderModuleState, state => state.customer);
export const getcustomers=createSelector(getCustomerState, fromCustomerReducer.getCustomers);
export const gettotalCustomerRecords=createSelector(getCustomerState, fromCustomerReducer.getTotalCustomerRecords);