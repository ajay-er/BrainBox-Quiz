import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'result',
    loadChildren: () =>
    import('../result-container/result-container.module').then(
      (m) => m.ResultContainerModule
      ),
    },
    {
      path: ':id',
      loadChildren: () =>
        import('../quiz-container/quiz-container.module').then(
          (m) => m.QuizContainerModule
        ),
    },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class QuizShellRoutingModule {}
