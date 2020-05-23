import {Action} from "@ngrx/store";
import {GetProduct} from "../../models/product/get-product";
import {ProductList} from "../../models/product/product-list";
import {Product} from "../../models/product/product";

export enum ProductActionTypes {
    LoadProducts = '[Product] Load Products',
    LoadProductsCompleted = '[Product] Load ProductsCompleted',
    GetProductById = '[Product] Get Product By Id',
    GetProductByIdCompleted = '[Product] Get Product By Id Complete',
    UpdateProduct='[Product] Update Product',
    UpdateProductCompleted='[Product] Update Product Completed',
    DeleteProduct='[Product] Delete Product',
    DeleteProductCompleted='[Product] Delete Product Completed',
}

export class LoadProducts implements Action {
    readonly type=ProductActionTypes.LoadProducts;

    constructor(public request: GetProduct){}
}

export class LoadProductsCompleted implements Action {
    readonly type=ProductActionTypes.LoadProductsCompleted;

    constructor(public payload: ProductList){}
}

export class GetProductById implements Action {
    readonly type=ProductActionTypes.GetProductById;

    constructor(public productId: number) {}
}

export class GetProductByIdCompleted implements Action {
    readonly type=ProductActionTypes.GetProductByIdCompleted;

    constructor(public payload: Product) {}
}

export class UpdateProduct implements Action {
    readonly type=ProductActionTypes.UpdateProduct;
    constructor(public request: Product) {}
}
export class UpdateProductCompleted implements Action {
    readonly type=ProductActionTypes.UpdateProductCompleted;
    constructor() {}
}

export class DeleteProduct implements Action {
    readonly type=ProductActionTypes.DeleteProduct;
    constructor(public productId: number) {}
}

export class DeleteProductCompleted implements Action {
    readonly type=ProductActionTypes.DeleteProductCompleted;
    constructor() {}
}

export type Actions=LoadProducts|LoadProductsCompleted|GetProductById|GetProductByIdCompleted|
    UpdateProduct|UpdateProductCompleted|DeleteProduct|DeleteProductCompleted;