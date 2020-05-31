import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomerPopupContainerComponent } from './customer-popup-container.component';

describe('CustomerPopupContainerComponent', () => {
  let component: CustomerPopupContainerComponent;
  let fixture: ComponentFixture<CustomerPopupContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CustomerPopupContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CustomerPopupContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
