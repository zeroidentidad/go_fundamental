import * as orderActions from '../actions/order.actions';
import {SearchOrderCriteria} from "../../models/order/search-order-criteria";
import {OrderListItem} from "../../models/order/order-list-item";

export interface State {
    query: SearchOrderCriteria;
    orders: OrderListItem[];
    totalRecords: number;
}

const initialState: State={
    query: new SearchOrderCriteria(),
    orders: [],
    totalRecords: 0,
}

export function OrderReducer(state=initialState, action: orderActions.Actions): State { 
    switch (action.type) {
        case orderActions.OrderActionTypes.UpdateOrderSearchCriteria: {
            if(action.payload) {
                return {
                    ...state,
                    query: action.payload
                }
            }
        }

        case orderActions.OrderActionTypes.LoadOrders: {
            return state;
        }

        case orderActions.OrderActionTypes.LoadOrdersCompleted: {
            return {
                ...state,
                orders: action.payload.data,
                totalRecords: action.payload.totalRecords
            }
        }        
    
        default:
            return state;
    }
}

export const getQuery=(state: State) => state.query;
export const getOrders=(state: State) => state.orders;
export const getTotalOrderRecords=(state: State) => state.totalRecords;