import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { HomeRoutingModule } from './home-routing.module';
import { HomeComponent } from './home.component';
import { BannerCarouselModule } from './ui/banner-carousel/banner-carousel.module';
import { CategoryCarouselModule } from './ui/category-carousel/category-carousel.module';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner'; 


@NgModule({
  declarations: [HomeComponent],
  imports: [
    CommonModule,
    HomeRoutingModule,
    BannerCarouselModule,
    CategoryCarouselModule,
    MatProgressSpinnerModule
  ],
})
export class HomeModule {}
