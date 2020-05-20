import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductCardListComponent } from './product-card-list.component';

describe('ProductCardListComponent', () => {
  let component: ProductCardListComponent;
  let fixture: ComponentFixture<ProductCardListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProductCardListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductCardListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
