import { TestBed } from '@angular/core/testing';

import { AppConfirmService } from './app-confirm.service';

describe('AppConfirmService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: AppConfirmService = TestBed.get(AppConfirmService);
    expect(service).toBeTruthy();
  });
});
