import {Action} from "@ngrx/store";
import {GetProduct} from "../../models/product/get-product";
import {ProductList} from "../../models/product/product-list";
import {Product} from "../../models/product/product";
import {ProductBestSeller} from "../../models/product/best-seller";

export enum ProductActionTypes {
    LoadProducts = '[Product] Load Products',
    LoadProductsCompleted = '[Product] Load ProductsCompleted',
    GetProductById = '[Product] Get Product By Id',
    GetProductByIdCompleted = '[Product] Get Product By Id Complete',
    UpdateProduct='[Product] Update Product',
    UpdateProductCompleted='[Product] Update Product Completed',
    DeleteProduct='[Product] Delete Product',
    DeleteProductCompleted='[Product] Delete Product Completed',
    AddProduct='[Product] Add Product',
    AddProductCompleted='[Product] Add Product Completed',
    GetBestSellers='[Product] Get Best Sellers',
    GetBestSellersCompleted='[Product] Get Best Sellers Completed'
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

export class AddProduct implements Action {
    readonly type=ProductActionTypes.AddProduct;
    constructor(public request: Product) {}
}

export class AddProductCompleted implements Action {
    readonly type=ProductActionTypes.AddProductCompleted;
    constructor() {}
}

export class GetBestSellers implements Action {
    readonly type=ProductActionTypes.GetBestSellers;
    constructor() {}
}

export class GetBestSellersCompleted implements Action {
    readonly type=ProductActionTypes.GetBestSellersCompleted;
    constructor(public payload: ProductBestSeller[]) {}
}

export type Actions=LoadProducts|LoadProductsCompleted|GetProductById|GetProductByIdCompleted|
    UpdateProduct|UpdateProductCompleted|DeleteProduct|DeleteProductCompleted|AddProduct|AddProductCompleted|GetBestSellers|GetBestSellersCompleted;