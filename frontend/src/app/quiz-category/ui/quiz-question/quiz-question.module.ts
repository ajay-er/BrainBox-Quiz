import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { QuizQuestionComponent } from './quiz-question.component';
import { LoadingButtonModule } from 'src/app/shared/ui/loading-button/loading-button.module';

@NgModule({
  declarations: [QuizQuestionComponent],
  imports: [CommonModule,LoadingButtonModule],
  exports:[QuizQuestionComponent]
})
export class QuizQuestionModule {}
