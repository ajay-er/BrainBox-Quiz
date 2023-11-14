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
    if(!this.isSvgValid){
      data.icon_svg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 text-red-600 h-6">
      <path fill-rule="evenodd" d="M2.25 12c0-5.385 4.365-9.75 9.75-9.75s9.75 4.365 9.75 9.75-4.365 9.75-9.75 9.75S2.25 17.385 2.25 12zM12 8.25a.75.75 0 01.75.75v3.75a.75.75 0 01-1.5 0V9a.75.75 0 01.75-.75zm0 8.25a.75.75 0 100-1.5.75.75 0 000 1.5z" clip-rule="evenodd" />
    </svg>`
    }    
    this.submissionData.emit(data);
    this.categoryForm.reset();
  }

  isSvgValid = false;
  validateSvg() {
    const parser = new DOMParser();
    const doc = parser.parseFromString(this.formData.iconTextarea, 'image/svg+xml');
    const isValidSvg = doc.documentElement.nodeName === 'svg';    
    this.santizedSvg = this.sanitizer.bypassSecurityTrustHtml(this.formData.iconTextarea)    
    this.isSvgValid = isValidSvg;
    
  }
  
}
