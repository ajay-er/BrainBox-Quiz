import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Output,
  ViewChild,
} from '@angular/core';
import { NgForm } from '@angular/forms';
import { DomSanitizer, SafeHtml } from '@angular/platform-browser';
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
  santizedSvg :SafeHtml = ''
  constructor(protected sanitizer: DomSanitizer) {}

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

  isSvgValid = true;
  validateSvg() {
    const parser = new DOMParser();
    const doc = parser.parseFromString(this.formData.iconTextarea, 'image/svg+xml');
    const isValidSvg = doc.documentElement.nodeName === 'svg';    
    this.santizedSvg = this.sanitizer.bypassSecurityTrustHtml(this.formData.iconTextarea)
    this.isSvgValid = isValidSvg;
  }
  
}
