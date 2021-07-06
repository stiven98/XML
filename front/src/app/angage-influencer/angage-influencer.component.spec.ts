import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AngageInfluencerComponent } from './angage-influencer.component';

describe('AngageInfluencerComponent', () => {
  let component: AngageInfluencerComponent;
  let fixture: ComponentFixture<AngageInfluencerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AngageInfluencerComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AngageInfluencerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
