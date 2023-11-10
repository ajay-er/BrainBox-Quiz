import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { CategoryRoutingModule } from './category-routing.module';
import { CategoryComponent } from './category.component';
import { TableModule } from 'src/app/shared/ui/table/table.module';
import { AddCategoryModule } from '../../ui/add-category/add-category.module';

@NgModule({
  declarations: [CategoryComponent],
  imports: [
    CommonModule,
    CategoryRoutingModule,
    TableModule,
    AddCategoryModule,
  ],
})
export class CategoryModule {}
