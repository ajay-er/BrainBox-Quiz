import { Component } from '@angular/core';
import { initCarousels, initFlowbite } from 'flowbite';

@Component({
  selector: 'app-banner-carousel',
  templateUrl: './banner-carousel.component.html',
  styleUrls: ['./banner-carousel.component.css'],
})
export class BannerCarouselComponent {
  ngOnInit() {
    initCarousels();
  }
}
