import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OrderMainContainerComponent } from './order-main-container.component';

describe('OrderMainContainerComponent', () => {
  let component: OrderMainContainerComponent;
  let fixture: ComponentFixture<OrderMainContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OrderMainContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OrderMainContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
