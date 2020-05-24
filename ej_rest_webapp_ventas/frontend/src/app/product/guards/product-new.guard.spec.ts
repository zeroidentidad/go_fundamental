import { TestBed, async, inject } from '@angular/core/testing';

import { ProductNewGuard } from './product-new.guard';

describe('ProductNewGuard', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ProductNewGuard]
    });
  });

  it('should ...', inject([ProductNewGuard], (guard: ProductNewGuard) => {
    expect(guard).toBeTruthy();
  }));
});
