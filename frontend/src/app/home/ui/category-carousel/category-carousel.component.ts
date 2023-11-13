import { Component, Input } from '@angular/core';
import { ICategory } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-category-carousel',
  templateUrl: './category-carousel.component.html',
  styleUrls: ['./category-carousel.component.css'],
})
export class CategoryCarouselComponent {
  @Input() categories: ICategory[] = [];
}
