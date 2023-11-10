import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddQuizContainerComponent } from './add-quiz-container.component';

describe('AddQuizContainerComponent', () => {
  let component: AddQuizContainerComponent;
  let fixture: ComponentFixture<AddQuizContainerComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AddQuizContainerComponent]
    });
    fixture = TestBed.createComponent(AddQuizContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
