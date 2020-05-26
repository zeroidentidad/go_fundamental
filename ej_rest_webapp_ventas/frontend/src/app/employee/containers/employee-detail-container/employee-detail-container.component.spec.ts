import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EmployeeDetailContainerComponent } from './employee-detail-container.component';

describe('EmployeeDetailContainerComponent', () => {
  let component: EmployeeDetailContainerComponent;
  let fixture: ComponentFixture<EmployeeDetailContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EmployeeDetailContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EmployeeDetailContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
