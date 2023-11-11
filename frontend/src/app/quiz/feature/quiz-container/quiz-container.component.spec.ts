import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QuizContainerComponent } from './quiz-container.component';

describe('QuizContainerComponent', () => {
  let component: QuizContainerComponent;
  let fixture: ComponentFixture<QuizContainerComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [QuizContainerComponent]
    });
    fixture = TestBed.createComponent(QuizContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
