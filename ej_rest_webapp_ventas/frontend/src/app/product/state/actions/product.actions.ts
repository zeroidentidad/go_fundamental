import {Action} from "@ngrx/store";
import {GetProduct} from "../../models/product/get-product";
import {ProductList} from "../../models/product/product-list";

export enum ProductActionTypes {
    LoadProducts = '[Product] Load Products',
    LoadProductsCompleted = '[Product] Load ProductsCompleted',
}

export class LoadProducts implements Action {
    readonly type=ProductActionTypes.LoadProducts;

    constructor(public request: GetProduct){}
}

export class LoadProductsCompleted implements Action {
    readonly type=ProductActionTypes.LoadProductsCompleted;

    constructor(public payload: ProductList){}
}

export type Actions=LoadProducts|LoadProductsCompleted;