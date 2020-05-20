import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductTableListComponent } from './product-table-list.component';

describe('ProductTableListComponent', () => {
  let component: ProductTableListComponent;
  let fixture: ComponentFixture<ProductTableListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ProductTableListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ProductTableListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
