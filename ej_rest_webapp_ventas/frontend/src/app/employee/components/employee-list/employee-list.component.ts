import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import {Employee} from "../../models/employee/employee";

@Component({
  selector: 'app-employee-list',
  templateUrl: './employee-list.component.html',
  styleUrls: ['./employee-list.component.scss']
})
export class EmployeeListComponent implements OnInit {

  @Input()
  employees: Employee[]=[];

  @Output()
  viewDetails: EventEmitter<any>=new EventEmitter<any>();

  @Output()
  delete: EventEmitter<any>=new EventEmitter<any>();

  @Output()
  edit: EventEmitter<number>=new EventEmitter<number>();

  constructor() { }

  ngOnInit() {
  }

  onViewDetails(event: any): void {
    this.viewDetails.emit(event)
  }

  onEdit(employeeId: number): void {
    this.edit.emit(employeeId)
  }

  onDelete(event: any): void {
    this.delete.emit(event);
  } 

}
