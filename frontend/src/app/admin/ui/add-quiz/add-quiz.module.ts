import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AddQuizComponent } from './add-quiz.component';
import { LoadingButtonModule } from 'src/app/shared/ui/loading-button/loading-button.module';

@NgModule({
  declarations: [AddQuizComponent],
  imports: [CommonModule,LoadingButtonModule],
  exports: [AddQuizComponent],
})
export class AddQuizModule {}
