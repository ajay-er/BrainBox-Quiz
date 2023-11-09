import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ResultContainerRoutingModule } from './result-container-routing.module';
import { ResultContainerComponent } from './result-container.component';
import { QuizResultModule } from '../../ui/quiz-result/quiz-result.module';

@NgModule({
  declarations: [ResultContainerComponent],
  imports: [CommonModule, ResultContainerRoutingModule,QuizResultModule],
})
export class ResultContainerModule {}
