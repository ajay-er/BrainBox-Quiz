import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { QuestionInputComponent } from './question-input.component';
import { ReactiveFormsModule ,FormGroup} from '@angular/forms';



@NgModule({
  declarations: [
    QuestionInputComponent
  ],
  imports: [
    CommonModule,ReactiveFormsModule
  ],exports:[QuestionInputComponent]
})
export class QuestionInputModule { }
