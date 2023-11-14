import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CategoryCarouselComponent } from './category-carousel.component';
import { RouterModule } from '@angular/router';

@NgModule({
  declarations: [CategoryCarouselComponent],
  imports: [CommonModule,RouterModule],
  exports: [CategoryCarouselComponent],
})
export class CategoryCarouselModule {}
