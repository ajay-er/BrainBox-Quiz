import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthContainerComponent } from './auth-container.component';
import { LoginComponent } from '../../ui/login/login.component';
import { SignupComponent } from '../../ui/signup/signup.component';

const routes: Routes = [
  {
    path: '',
    component: AuthContainerComponent,
    children: [
      { path: 'login', component: LoginComponent },
      { path: 'signup', component: SignupComponent },
      { path: '', redirectTo: 'login', pathMatch: 'full' },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AuthContainerRoutingModule {}
