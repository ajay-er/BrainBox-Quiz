import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'sports',
    pathMatch: 'full',
  },
  {
    path: ':category',
    loadChildren: () =>
      import('../quiz-category/quiz-category.module').then(
        (m) => m.QuizCategoryModule
      ),
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class QuizCategoryShellRoutingModule {}
