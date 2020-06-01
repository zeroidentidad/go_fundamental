import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductPopupContainerComponent } from './product-popup-container.component';

describe('ProductPopupContainerComponent', () => {
  let component: ProductPopupContainerComponent;
  let fixture: ComponentFixture<ProductPopupContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProductPopupContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductPopupContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
