import { Component, OnInit, Input } from '@angular/core';
import {BestEmployee} from "../../models/employee/best-employee";

@Component({
  selector: 'app-best-employee',
  templateUrl: './best-employee.component.html',
  styleUrls: ['./best-employee.component.scss']
})
export class BestEmployeeComponent implements OnInit {

@Input()
employee: BestEmployee;

  constructor() { }

  ngOnInit() {
  }

}
