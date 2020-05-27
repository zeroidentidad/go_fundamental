import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BestEmployeeComponent } from './best-employee.component';

describe('BestEmployeeComponent', () => {
  let component: BestEmployeeComponent;
  let fixture: ComponentFixture<BestEmployeeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BestEmployeeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BestEmployeeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
