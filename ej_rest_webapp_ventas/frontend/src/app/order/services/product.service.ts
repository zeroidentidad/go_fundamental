import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {GetProducts} from "../models/product/get-product";
import {Observable} from "rxjs";
import {ProductList} from "src/app/product/models/product/product-list";
import {environment} from "src/environments/environment";

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  constructor(private httpClient: HttpClient) { }

  getProducts(request: GetProducts): Observable<ProductList> {
    return this.httpClient.post<ProductList>(`${environment.ApiURL}products/paginated`, request);
  }

}
