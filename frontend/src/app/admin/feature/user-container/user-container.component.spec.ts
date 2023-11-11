import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserContainerComponent } from './user-container.component';

describe('UserContainerComponent', () => {
  let component: UserContainerComponent;
  let fixture: ComponentFixture<UserContainerComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [UserContainerComponent]
    });
    fixture = TestBed.createComponent(UserContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
