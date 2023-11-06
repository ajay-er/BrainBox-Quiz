import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SignupComponent } from './signup.component';
import { LoadingButtonModule } from 'src/app/shared/ui/loading-button/loading-button.module';
import { ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

@NgModule({
  declarations: [SignupComponent],
  imports: [
    CommonModule,
    RouterModule,
    LoadingButtonModule,
    ReactiveFormsModule,
  ],
  exports: [SignupComponent],
})
export class SignupModule {}
