import * as orderActions from '../actions/order.actions';
import {SearchOrderCriteria} from "../../models/order/search-order-criteria";
import {OrderListItem} from "../../models/order/order-list-item";

export interface State {
    query: SearchOrderCriteria;
}

const initialState: State={
    query: new SearchOrderCriteria(),
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
    
        default:
            return state;
    }
}

export const getQuery=(state: State) => state.query;
