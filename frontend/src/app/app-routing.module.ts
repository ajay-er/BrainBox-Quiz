import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NotFoundComponent } from './shared/ui/not-found/not-found.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'home',
    pathMatch: 'full',
  },
  {
    path: 'home',
    loadChildren: () => import('./home/home.module').then((m) => m.HomeModule),
  },
  {
    path: 'auth',
    loadChildren: () =>
      import('./auth/feature/auth-shell/auth-shell.module').then(
        (m) => m.AuthShellModule
      ),
    // canActivate: [unauthenticatedGuard],
  },
  {
    path: 'admin',
    loadChildren: () =>
      import('./admin/feature/admin-shell/admin-shell.module').then(
        (m) => m.AdminShellModule
      ),
    // canActivate: [unauthenticatedGuard],
  },
  {
    path: 'category',
    loadChildren: () =>
      import(
        './quiz-category/feature/quiz-category-shell/quiz-category-shell.module'
      ).then((m) => m.QuizCategoryShellModule),
  },
  {
    path: 'quiz',
    loadChildren: () =>
      import('./quiz/feature/quiz-shell/quiz-shell.module').then(
        (m) => m.QuizShellModule
      ),
  },
  {
    path: '**',
    component: NotFoundComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
