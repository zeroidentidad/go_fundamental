import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { OrderManageContainerComponent } from './order-manage-container.component';

describe('OrderManageContainerComponent', () => {
  let component: OrderManageContainerComponent;
  let fixture: ComponentFixture<OrderManageContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ OrderManageContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OrderManageContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
