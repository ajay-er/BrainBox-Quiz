import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AuthContainerComponent } from './auth-container.component';

describe('AuthContainerComponent', () => {
  let component: AuthContainerComponent;
  let fixture: ComponentFixture<AuthContainerComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AuthContainerComponent]
    });
    fixture = TestBed.createComponent(AuthContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
