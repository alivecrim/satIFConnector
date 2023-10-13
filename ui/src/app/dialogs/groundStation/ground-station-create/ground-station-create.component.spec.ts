import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroundStationCreateComponent } from './ground-station-create.component';

describe('GroundStationCreateComponent', () => {
  let component: GroundStationCreateComponent;
  let fixture: ComponentFixture<GroundStationCreateComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [GroundStationCreateComponent]
    });
    fixture = TestBed.createComponent(GroundStationCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
