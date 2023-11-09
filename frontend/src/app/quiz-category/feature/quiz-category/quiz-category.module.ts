import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { QuizCategoryRoutingModule } from './quiz-category-routing.module';
import { QuizCategoryComponent } from './quiz-category.component';
import { QuizQuestionModule } from '../../ui/quiz-question/quiz-question.module';

@NgModule({
  declarations: [QuizCategoryComponent],
  imports: [CommonModule, QuizCategoryRoutingModule, QuizQuestionModule],
})
export class QuizCategoryModule {}
