import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StoryArchiveComponent } from './story-archive.component';

describe('StoryArchiveComponent', () => {
  let component: StoryArchiveComponent;
  let fixture: ComponentFixture<StoryArchiveComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ StoryArchiveComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(StoryArchiveComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
