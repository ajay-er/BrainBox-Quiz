import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditUserInputComponent } from './edit-user-input.component';

describe('EditUserInputComponent', () => {
  let component: EditUserInputComponent;
  let fixture: ComponentFixture<EditUserInputComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [EditUserInputComponent]
    });
    fixture = TestBed.createComponent(EditUserInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
