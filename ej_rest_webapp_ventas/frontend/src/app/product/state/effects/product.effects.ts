import {Injectable} from "@angular/core";
import {ProductService} from "../../services/product.service";
import {Actions, Effect, ofType} from "@ngrx/effects";
import * as productActions from "../actions/product.actions";
import {switchMap, map} from "rxjs/operators";
import {Router} from "@angular/router";

@Injectable()
export class ProductEffects{
    constructor(private productService: ProductService, private actions$:Actions, private router: Router){}

    @Effect()
    getProducts$ = this.actions$.pipe(
        ofType<productActions.LoadProducts>(productActions.ProductActionTypes.LoadProducts),
        switchMap(action=>this.productService.getProducts(action.request)
        .pipe(
            map(data=>new productActions.LoadProductsCompleted(data))
        ))
    )

    @Effect()
    getProductById$ = this.actions$.pipe(
        ofType<productActions.GetProductById>(productActions.ProductActionTypes.GetProductById),
        switchMap(action => this.productService.getProductById(action.productId)
        .pipe(
            map(data => new productActions.GetProductByIdCompleted(data))
        ))        
    )

    @Effect()
    updateProduct$=this.actions$.pipe(
        ofType<productActions.UpdateProduct>(productActions.ProductActionTypes.UpdateProduct),
        switchMap(action => this.productService.updateProduct(action.request)
            .pipe(
                map(_ => {
                    this.router.navigate(['/product/list'])
                    return new productActions.UpdateProductCompleted()
                }
                )))
    ); 
    
    @Effect()
    deleteProduct$=this.actions$.pipe(
        ofType<productActions.DeleteProduct>(productActions.ProductActionTypes.DeleteProduct),
        switchMap(action => this.productService.deleteProduct(action.productId)
            .pipe(
                map(_ => {
                    return new productActions.DeleteProductCompleted();
                })
            ))
    );
}