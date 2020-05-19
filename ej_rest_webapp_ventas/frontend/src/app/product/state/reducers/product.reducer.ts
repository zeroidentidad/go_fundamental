import {Product} from "../../models/product/product";
import * as ProductActions from "../actions/product.actions";

export interface State {
    products: Product[];
    totalRecords: number;
}

const initialState: State = {
    products: [],
    totalRecords: 0
}

export function ProductReducer(state=initialState, action: ProductActions.Actions) {
    switch (action.type) {
        case ProductActions.ProductActionTypes.LoadProducts:
            return state;

        case ProductActions.ProductActionTypes.LoadProductsCompleted:
            return {
                ...state, 
                products: action.payload.data,
                totalRecords: action.payload.totalRecords
            }
    
        default:
            return state;
    }
}

export const getProducts = (state:State)=>state.products;
export const getTotalRecords = (state:State)=>state.totalRecords;