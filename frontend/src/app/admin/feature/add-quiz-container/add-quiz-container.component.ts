import { HttpClient } from '@angular/common/http';
import { Component, inject } from '@angular/core';
import { AdminService } from '../../data-access/admin.service';
import { ICategory } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-add-quiz-container',
  templateUrl: './add-quiz-container.component.html',
  styleUrls: ['./add-quiz-container.component.css']
})
export class AddQuizContainerComponent {
  private adminService = inject(AdminService)
  categories:ICategory[] = []
  ngOnInit(){
    this.adminService.getAllCategories().subscribe((res:any)=>{
      console.log(res);
      if (res.data?.Categories) {
        this.categories = res.data.Categories;
      }
    })
  }
}
