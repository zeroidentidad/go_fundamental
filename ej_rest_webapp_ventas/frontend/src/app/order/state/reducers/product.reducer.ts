import * as productActions from '../actions/product.actions';
import {Product} from "../../models/product/product";

export interface State {
    products: Product[];
    totalProductRecords: number;
}
const initialState: State={
    products: [],
    totalProductRecords: 0
};

export function ProductReducer(state=initialState, action: productActions.Actions): State {
    switch(action.type) {
        case productActions.ProductActionTypes.LoadProducts: {
            return state;
        }
        case productActions.ProductActionTypes.LoadProductsCompleted: {
            return {
                ...state,
                products: action.payload.data,
                totalProductRecords: action.payload.totalRecords
            };
        }
        default: return state;
    }
}

export const getProducts=(state: State) => state.products;
export const getTotalProductRecords=(state: State) => state.totalProductRecords;