import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AddQuizContainerRoutingModule } from './add-quiz-container-routing.module';
import { AddQuizContainerComponent } from './add-quiz-container.component';
import { AddQuizModule } from '../../ui/add-quiz/add-quiz.module';

@NgModule({
  declarations: [AddQuizContainerComponent],
  imports: [CommonModule, AddQuizContainerRoutingModule, AddQuizModule,
  ],
})
export class AddQuizContainerModule {}
