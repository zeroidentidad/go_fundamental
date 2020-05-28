import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EmployeeEditContainerComponent } from './employee-edit-container.component';

describe('EmployeeEditContainerComponent', () => {
  let component: EmployeeEditContainerComponent;
  let fixture: ComponentFixture<EmployeeEditContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EmployeeEditContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EmployeeEditContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
