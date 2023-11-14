import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QuestionInputComponent } from './question-input.component';

describe('QuestionInputComponent', () => {
  let component: QuestionInputComponent;
  let fixture: ComponentFixture<QuestionInputComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [QuestionInputComponent]
    });
    fixture = TestBed.createComponent(QuestionInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
