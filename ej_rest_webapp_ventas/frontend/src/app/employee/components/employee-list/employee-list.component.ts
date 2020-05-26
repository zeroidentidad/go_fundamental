import { Component, OnInit, Input } from '@angular/core';
import {Employee} from "../../models/employee/employee";

@Component({
  selector: 'app-employee-list',
  templateUrl: './employee-list.component.html',
  styleUrls: ['./employee-list.component.scss']
})
export class EmployeeListComponent implements OnInit {

  @Input()
  employees: Employee[]=[];

  constructor() { }

  ngOnInit() {
  }

}
