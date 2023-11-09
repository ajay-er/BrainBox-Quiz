import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BannerCarouselComponent } from './banner-carousel.component';

@NgModule({
  declarations: [BannerCarouselComponent],
  imports: [CommonModule],
  exports: [BannerCarouselComponent],
})
export class BannerCarouselModule {}
