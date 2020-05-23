import {Product} from "../../models/product/product";
import * as ProductActions from "../actions/product.actions";

export interface State {
    products: Product[];
    totalRecords: number;
    product: Product;
}

const initialState: State = {
    products: [],
    totalRecords: 0,
    product: new Product()
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

        case ProductActions.ProductActionTypes.GetProductById:
            return state;

        case ProductActions.ProductActionTypes.GetProductByIdCompleted:
            return {
                ...state,
                product: action.payload
            }  
            
        case ProductActions.ProductActionTypes.UpdateProduct: {
            return state;
        }
        
        case ProductActions.ProductActionTypes.UpdateProductCompleted: {
            return state;
        }
        
        case ProductActions.ProductActionTypes.DeleteProduct: {
            return state;
        }

        case ProductActions.ProductActionTypes.DeleteProductCompleted: {
            return state;
        }        
    
        default:
            return state;
    }
}

export const getProducts = (state:State)=>state.products;
export const getTotalRecords = (state:State)=>state.totalRecords;
export const getProduct = (state:State)=>state.product;