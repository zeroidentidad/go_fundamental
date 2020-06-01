import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomerPopupComponent } from './customer-popup.component';

describe('CustomerPopupComponent', () => {
  let component: CustomerPopupComponent;
  let fixture: ComponentFixture<CustomerPopupComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CustomerPopupComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CustomerPopupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
