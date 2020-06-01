import {Action} from '@ngrx/store';
import {GetProducts} from "../../models/product/get-product";
import {ProductList} from "../../models/product/product-list";

export enum ProductActionTypes {
    LoadProducts='[Order-Product] Load Products',
    LoadProductsCompleted='[Order-Product] Load Products Completed'
}

export class LoadProducts implements Action {
    readonly type=ProductActionTypes.LoadProducts;
    constructor(public request: GetProducts) {}
}

export class LoadProductsCompleted implements Action {
    readonly type=ProductActionTypes.LoadProductsCompleted;
    constructor(public payload: ProductList) {}
}

export type Actions=|LoadProducts|LoadProductsCompleted;