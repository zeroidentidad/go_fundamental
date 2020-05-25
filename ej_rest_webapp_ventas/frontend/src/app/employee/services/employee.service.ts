import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {GetEmployee} from "../models/employee/get-employee";
import {EmployeeList} from "../models/employee/employee-list";
import {Observable} from "rxjs";
import {environment} from "src/environments/environment";
import {Employee} from "../models/employee/employee";
import {map} from "rxjs/operators";
import {BestEmployee} from "../models/employee/best-employee";

@Injectable({
  providedIn: 'root'
})
export class EmployeeService {

  constructor(private httpClient:HttpClient) { }

  getEmployees(request: GetEmployee):Observable<EmployeeList>{
    return this.httpClient.post<EmployeeList>(`${environment.ApiURL}employees/paginated`, request)
  }

  getEmployeeById(id: number): Observable<Employee> {
    return this.httpClient.get<Employee>(`${environment.ApiURL}employees/${id}`);
  }  

  createEmployee(request: Employee): Observable<Response> {
    return this.httpClient.post(`${environment.ApiURL}employees`, request)
      .pipe(
        map((response: Response) => response)
      );
  }   

  updateEmployee(request: Employee): Observable<Response> {
    return this.httpClient.put(`${environment.ApiURL}employees`, request)
      .pipe(
        map((response: Response) => response)
      );
  } 
  
  deleteEmployee(id: number): Observable<Response> {
    return this.httpClient.delete(`${environment.ApiURL}employees/${id}`)
      .pipe(
        map((response: Response) => response)
      );
  }
  
  getBestEmployee(): Observable<BestEmployee> {
    return this.httpClient.get<BestEmployee>(`${environment.ApiURL}employees/best`);
  } 

}
