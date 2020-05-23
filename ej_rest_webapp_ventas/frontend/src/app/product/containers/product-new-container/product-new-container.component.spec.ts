import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductNewContainerComponent } from './product-new-container.component';

describe('ProductNewContainerComponent', () => {
  let component: ProductNewContainerComponent;
  let fixture: ComponentFixture<ProductNewContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProductNewContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductNewContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
