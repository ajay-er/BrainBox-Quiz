import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'users',
    pathMatch: 'full',
  },
  {
    path: 'add',
    loadChildren: () =>
      import('../add-quiz-container/add-quiz-container.module').then(
        (m) => m.AddQuizContainerModule
      ),
  },
  {
    path: 'edit',
    loadChildren: () =>
      import('../edit-user/edit-user.module').then((m) => m.EditUserModule),
  },
  {
    path: 'users',
    loadChildren: () =>
      import('../user-container/user-container.module').then(
        (m) => m.UserContainerModule
      ),
  },
  {
    path: 'category',
    loadChildren: () =>
      import('../category/category.module').then((m) => m.CategoryModule),
  },
  {
    path: 'quiz',
    loadChildren: () =>
      import('../add-quiz-container/add-quiz-container.module').then(
        (m) => m.AddQuizContainerModule
      ),
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AdminShellRoutingModule {}
