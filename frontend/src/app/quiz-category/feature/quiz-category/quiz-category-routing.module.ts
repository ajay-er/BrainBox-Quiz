import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { QuizCategoryComponent } from './quiz-category.component';

const routes: Routes = [
  {
    path: '',
    component: QuizCategoryComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class QuizCategoryRoutingModule {}
