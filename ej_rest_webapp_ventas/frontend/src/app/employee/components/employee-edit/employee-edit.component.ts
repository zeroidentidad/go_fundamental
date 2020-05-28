import { Component, OnInit, Input, Output, EventEmitter, SimpleChanges } from '@angular/core';
import {Employee} from "../../models/employee/employee";
import {FormGroup, FormBuilder, Validators} from "@angular/forms";
import {validatePhone} from "src/app/shared/validators/phone.validator";

@Component({
  selector: 'app-employee-edit',
  templateUrl: './employee-edit.component.html',
  styleUrls: ['./employee-edit.component.scss']
})
export class EmployeeEditComponent implements OnInit {

  @Input()
  employee: Employee;

  @Output()
  edit: EventEmitter<Employee>=new EventEmitter<Employee>();

  employeeEditForm: FormGroup;

  constructor(private fb: FormBuilder) {

  }

  ngOnChanges(changes: SimpleChanges): void {
    if(changes.employee&&changes.employee.currentValue) {
      this.buildEmployeeEditForm()
    }
  }

  ngOnInit() {
  }

  buildEmployeeEditForm(): void {
    this.employeeEditForm=this.fb.group({
      firstName: [this.employee.firstName, [Validators.required]],
      lastName: [this.employee.lastName, [Validators.required]],
      company: [this.employee.company, []],
      address: [this.employee.address, [Validators.required]],
      businessPhone: [this.employee.businessPhone, [validatePhone]],
      emailAddress: [this.employee.emailAddress, [Validators.required, Validators.email]],
      faxNumber: [this.employee.faxNumber, []],
      homePhone: [this.employee.homePhone, []],
      jobTitle: [this.employee.jobTitle, []],
      mobilePhone: [this.employee.mobilePhone, [validatePhone]],
    });
  }

  onEdit(): void {
    if(this.employeeEditForm.valid) {
      if(this.employeeEditForm.dirty) {
        const editedEmployee={... this.employee, ...this.employeeEditForm.value}
        this.edit.emit(editedEmployee)
      }
    }
  }  

}
