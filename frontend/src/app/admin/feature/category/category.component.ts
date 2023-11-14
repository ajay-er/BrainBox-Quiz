import { Component, inject } from '@angular/core';
import { ICategory } from 'src/app/shared/interfaces';
import { AdminService } from '../../data-access/admin.service';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css'],
})
export class CategoryComponent {
  private adminService = inject(AdminService);
  submitCategoryData(data: ICategory) {    
    this.adminService.addNewCategory(data).subscribe((res: any) => {
      console.log(res);
    });
  }
}
