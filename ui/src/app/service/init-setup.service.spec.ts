import { TestBed } from '@angular/core/testing';

import { InitSetupService } from './init-setup.service';

describe('InitSetupService', () => {
  let service: InitSetupService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(InitSetupService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
