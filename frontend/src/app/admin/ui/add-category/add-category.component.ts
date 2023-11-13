import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Output,
  ViewChild,
} from '@angular/core';
import { NgForm } from '@angular/forms';
import { ICategory } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-add-category',
  templateUrl: './add-category.component.html',
  styleUrls: ['./add-category.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AddCategoryComponent {
  @ViewChild('categoryForm') categoryForm!: NgForm;
  @Output() submissionData: EventEmitter<any> = new EventEmitter();

  formData = {
    categoryName: '',
    iconTextarea: '',
  };

  onSubmit(formData: any) {
    const data: ICategory = {
      category_name: formData.categoryName,
      icon_svg: formData.iconTextarea,
    };
    this.submissionData.emit(data);
    this.categoryForm.reset();
  }
}
