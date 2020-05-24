import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BestSellerProductContainerComponent } from './best-seller-product-container.component';

describe('BestSellerProductContainerComponent', () => {
  let component: BestSellerProductContainerComponent;
  let fixture: ComponentFixture<BestSellerProductContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BestSellerProductContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BestSellerProductContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
