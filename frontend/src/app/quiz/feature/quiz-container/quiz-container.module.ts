import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { QuizContainerRoutingModule } from './quiz-container-routing.module';
import { QuizContainerComponent } from './quiz-container.component';
import { QuizCardModule } from '../../ui/quiz-card/quiz-card.module';

@NgModule({
  declarations: [QuizContainerComponent],
  imports: [CommonModule, QuizContainerRoutingModule,QuizCardModule],
})
export class QuizContainerModule {}
