import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserContainerComponent } from './user-container.component';

const routes: Routes = [
  {
    path: '',
    component: UserContainerComponent,
  },
  {
    path: 'edit',
    loadChildren: () =>
      import('../edit-user/edit-user.module').then((m) => m.EditUserModule),
  },
  {
    path: 'add',
    loadChildren: () =>
      import('../add-user/add-user.module').then((m) => m.AddUserModule),
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class UserContainerRoutingModule {}
