import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {GetProduct} from "../models/product/get-product";
import {Observable} from "rxjs";
import {ProductList} from "../models/product/product-list";
import {environment} from "src/environments/environment";
import {Product} from "../models/product/product";

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
}
