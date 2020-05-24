import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {GetProduct} from "../models/product/get-product";
import {Observable} from "rxjs";
import {ProductList} from "../models/product/product-list";
import {environment} from "src/environments/environment";
import {Product} from "../models/product/product";
import {map} from "rxjs/operators";
import {ProductBestSeller} from "../models/product/best-seller";

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  constructor(private httpClient:HttpClient) { }

  getProducts(request:GetProduct):Observable<ProductList>{
    return this.httpClient.post<ProductList>(`${environment.ApiURL}products/paginated`, request);
  }

  getProductById(id:number):Observable<Product>{
    return this.httpClient.get<Product>(`${environment.ApiURL}products/${id}`);
  }

  updateProduct(product:Product):Observable<Response>{
    return this.httpClient.put(`${environment.ApiURL}products`, product)
    .pipe(
      map((response: Response) => response)
    )
  }

  deleteProduct(id: number): Observable<Response> {
    return this.httpClient.delete(`${environment.ApiURL}products/${id}`)
      .pipe(
        map((response: Response) => response)
      );
  }  

  addProduct(product: Product): Observable<Response> {
    return this.httpClient.post(`${environment.ApiURL}products`, product)
      .pipe(
        map((response: Response) => response)
      );
  }

  getBestSellers(): Observable<ProductBestSeller[]> {
    return this.httpClient.get(`${environment.ApiURL}products/bestSellers`)
      .pipe(
        map((response: any) => {
          return response.data.map((product: any) => {
            return ProductBestSeller.mapFromResponse(product, response.totalVentas)
          })
        })
      );
  }

}
