import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SatelliteCreateComponent } from './satellite-create.component';

describe('SatelliteCreateComponent', () => {
  let component: SatelliteCreateComponent;
  let fixture: ComponentFixture<SatelliteCreateComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SatelliteCreateComponent]
    });
    fixture = TestBed.createComponent(SatelliteCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
