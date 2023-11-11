import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { QuizResultComponent } from './quiz-result.component';
import { LoadingButtonModule } from 'src/app/shared/ui/loading-button/loading-button.module';

@NgModule({
  declarations: [QuizResultComponent],
  imports: [CommonModule,LoadingButtonModule],
  exports: [QuizResultComponent],
})
export class QuizResultModule {}
