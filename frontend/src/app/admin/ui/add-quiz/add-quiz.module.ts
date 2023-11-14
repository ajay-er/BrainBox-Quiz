import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AddQuizComponent } from './add-quiz.component';
import { LoadingButtonModule } from 'src/app/shared/ui/loading-button/loading-button.module';
import { ReactiveFormsModule } from '@angular/forms';
import { QuestionInputModule } from '../question-input/question-input.module';

@NgModule({
  declarations: [AddQuizComponent],
  imports: [
    CommonModule,
    LoadingButtonModule,
    ReactiveFormsModule,
    QuestionInputModule,
  ],
  exports: [AddQuizComponent],
})
export class AddQuizModule {}
