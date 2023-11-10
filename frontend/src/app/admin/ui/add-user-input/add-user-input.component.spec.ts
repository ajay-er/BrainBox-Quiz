import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddUserInputComponent } from './add-user-input.component';

describe('AddUserInputComponent', () => {
  let component: AddUserInputComponent;
  let fixture: ComponentFixture<AddUserInputComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AddUserInputComponent]
    });
    fixture = TestBed.createComponent(AddUserInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
