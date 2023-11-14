import { Component, Input } from '@angular/core';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
import { ICategory } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-category-carousel',
  templateUrl: './category-carousel.component.html',
  styleUrls: ['./category-carousel.component.css'],
})
export class CategoryCarouselComponent {
  @Input() categories: ICategory[] = [];
  
  constructor(private sanitizer: DomSanitizer) {}
  
  sanitizeHtml(html: string): SafeHtml {
    return this.sanitizer.bypassSecurityTrustHtml(html);
  }
  

}
