import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AddCategoryComponent } from './add-category.component';
import { FormsModule } from '@angular/forms';
import { LoadingButtonModule } from 'src/app/shared/ui/loading-button/loading-button.module';

@NgModule({
  declarations: [AddCategoryComponent],
  imports: [CommonModule, FormsModule, LoadingButtonModule],
  exports: [AddCategoryComponent],
})
export class AddCategoryModule {}
