import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthContainerComponent } from './auth-container.component';
import { LoginComponent } from '../../ui/login/login.component';
import { SignupComponent } from '../../ui/signup/signup.component';
import { adminUnauthGuard } from 'src/app/shared/guards/admin-unauth.guard';
import { unauthGuard } from 'src/app/shared/guards/unauth.guard';

const routes: Routes = [
  {
    path: '',
    component: AuthContainerComponent,
    children: [
      { path: 'login', component: LoginComponent, canActivate: [unauthGuard] },
      {
        path: 'signup',
        component: SignupComponent,
        canActivate: [unauthGuard],
      },
      {
        path: 'admin-login',
        component: LoginComponent,
        canActivate: [adminUnauthGuard],
      },
      { path: '', redirectTo: 'login', pathMatch: 'full' },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AuthContainerRoutingModule {}
