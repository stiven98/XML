import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LikedDislikedPostsComponent } from './liked-disliked-posts.component';

describe('LikedDislikedPostsComponent', () => {
  let component: LikedDislikedPostsComponent;
  let fixture: ComponentFixture<LikedDislikedPostsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LikedDislikedPostsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LikedDislikedPostsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
