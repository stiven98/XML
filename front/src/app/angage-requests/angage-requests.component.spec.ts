import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AngageRequestsComponent } from './angage-requests.component';

describe('AngageRequestsComponent', () => {
  let component: AngageRequestsComponent;
  let fixture: ComponentFixture<AngageRequestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AngageRequestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AngageRequestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
