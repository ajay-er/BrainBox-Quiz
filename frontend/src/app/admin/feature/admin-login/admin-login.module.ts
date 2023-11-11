import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AdminLoginRoutingModule } from './admin-login-routing.module';
import { AdminLoginComponent } from './admin-login.component';
import { LoginModule } from '../../ui/login/login.module';

@NgModule({
  declarations: [AdminLoginComponent],
  imports: [CommonModule, AdminLoginRoutingModule, LoginModule],
})
export class AdminLoginModule {}
