import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AgentRequestsComponent } from './agent-requests.component';

describe('AgentRequestsComponent', () => {
  let component: AgentRequestsComponent;
  let fixture: ComponentFixture<AgentRequestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AgentRequestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AgentRequestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
