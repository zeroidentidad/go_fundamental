import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EmployeeMainContainerComponent } from './employee-main-container.component';

describe('EmployeeMainContainerComponent', () => {
  let component: EmployeeMainContainerComponent;
  let fixture: ComponentFixture<EmployeeMainContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EmployeeMainContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EmployeeMainContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
