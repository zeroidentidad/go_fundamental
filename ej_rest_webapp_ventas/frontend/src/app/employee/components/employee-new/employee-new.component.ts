import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import {Employee} from "../../models/employee/employee";
import {FormGroup, FormBuilder, Validators} from "@angular/forms";
import {validatePhone} from "src/app/shared/validators/phone.validator";

@Component({
  selector: 'app-employee-new',
  templateUrl: './employee-new.component.html',
  styleUrls: ['./employee-new.component.scss']
})
export class EmployeeNewComponent implements OnInit {

  @Output()
  add: EventEmitter<Employee>=new EventEmitter<Employee>();

  employeeForm: FormGroup;
  employee: Employee;

  constructor(private formBuilder: FormBuilder) { 
    this.buildNewForm();
  }

  ngOnInit() {
  }

  buildNewForm(): void {
    this.employeeForm=this.formBuilder.group({
      firstName: ['', [Validators.required]],
      lastName: ['', [Validators.required]],
      company: ['', []],
      address: ['', [Validators.required]],
      businessPhone: ['', [validatePhone]],
      emailAddress: ['', [Validators.required, Validators.email]],
      faxNumber: ['', []],
      homePhone: ['', []],
      jobTitle: ['', []],
      mobilePhone: ['', [validatePhone]],
    });
  }  

  onCreate(): void {
    if(this.employeeForm.valid) {
      if(this.employeeForm.dirty) {
        const newEmployee={...this.employee, ...this.employeeForm.value};
        this.add.emit(newEmployee);
      } else {
        console.log("no se actualizo nada")
      }
    }
  }  

}
