import {ActionReducerMap, createFeatureSelector, createSelector} from '@ngrx/store';
import * as fromOrderReducer from '../reducers/order.reducer';

export interface OrderState {
    order: fromOrderReducer.State;
}

export const reducers: ActionReducerMap<OrderState>={
    order: fromOrderReducer.OrderReducer,
}

export const getOrderModuleState=createFeatureSelector<OrderState>('order');

/// Order State
export const getOrderState=createSelector(getOrderModuleState, state => state.order);
export const getQuery=createSelector(getOrderState, fromOrderReducer.getQuery);
export const getOrders=createSelector(getOrderState, fromOrderReducer.getOrders);
export const getTotalRecords=createSelector(getOrderState, fromOrderReducer.getTotalOrderRecords);