import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UpsertInputComponent } from './upsert-input.component';
import { LoadingButtonModule } from 'src/app/shared/ui/loading-button/loading-button.module';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [UpsertInputComponent],
  imports: [CommonModule, LoadingButtonModule, ReactiveFormsModule],
  exports: [UpsertInputComponent],
})
export class UpsertInputModule {}
