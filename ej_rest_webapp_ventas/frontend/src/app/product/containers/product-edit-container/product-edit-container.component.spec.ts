import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductEditContainerComponent } from './product-edit-container.component';

describe('ProductEditContainerComponent', () => {
  let component: ProductEditContainerComponent;
  let fixture: ComponentFixture<ProductEditContainerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProductEditContainerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductEditContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
