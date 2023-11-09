import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { QuizContainerRoutingModule } from './quiz-container-routing.module';
import { QuizContainerComponent } from './quiz-container.component';

@NgModule({
  declarations: [QuizContainerComponent],
  imports: [CommonModule, QuizContainerRoutingModule],
})
export class QuizContainerModule {}
