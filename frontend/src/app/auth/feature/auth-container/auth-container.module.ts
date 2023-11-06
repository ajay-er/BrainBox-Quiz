import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AuthContainerRoutingModule } from './auth-container-routing.module';
import { AuthContainerComponent } from './auth-container.component';
import { LoginModule } from '../../ui/login/login.module';
import { LoadingButtonModule } from 'src/app/shared/ui/loading-button/loading-button.module';

@NgModule({
  declarations: [AuthContainerComponent],
  imports: [
    CommonModule,
    AuthContainerRoutingModule,
    LoginModule,
    LoadingButtonModule,
  ],
})
export class AuthContainerModule {}
