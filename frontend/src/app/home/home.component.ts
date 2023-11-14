import { Component, inject } from '@angular/core';
import { ICategory } from '../shared/interfaces';
import { HomeService } from './data-access/home.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
})
export class HomeComponent {
  isLoading: boolean = false;
  categoriesData: ICategory[] = [];

  private homeService = inject(HomeService);

  ngOnInit() {
    this.homeService.getAllCategories().subscribe((res: any) => {
      if (res.data?.Categories) {
        this.categoriesData = res.data.Categories;
      }
    });
  }
}
