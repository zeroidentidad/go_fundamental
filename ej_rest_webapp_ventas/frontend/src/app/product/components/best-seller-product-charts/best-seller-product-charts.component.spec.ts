import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BestSellerProductChartsComponent } from './best-seller-product-charts.component';

describe('BestSellerProductChartsComponent', () => {
  let component: BestSellerProductChartsComponent;
  let fixture: ComponentFixture<BestSellerProductChartsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BestSellerProductChartsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BestSellerProductChartsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
