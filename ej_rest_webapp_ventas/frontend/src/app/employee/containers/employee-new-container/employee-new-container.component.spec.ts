import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EmployeeNewContainerComponent } from './employee-new-container.component';

describe('EmployeeNewContainerComponent', () => {
  let component: EmployeeNewContainerComponent;
  let fixture: ComponentFixture<EmployeeNewContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EmployeeNewContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EmployeeNewContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
