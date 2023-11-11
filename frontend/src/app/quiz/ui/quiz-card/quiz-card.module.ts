import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { QuizCardComponent } from './quiz-card.component';

@NgModule({
  declarations: [QuizCardComponent],
  imports: [CommonModule],
  exports: [QuizCardComponent],
})
export class QuizCardModule {}
