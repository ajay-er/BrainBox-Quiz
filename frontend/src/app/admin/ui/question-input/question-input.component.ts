import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Input,
  Output,
} from '@angular/core';
import {
  FormGroup,
  FormArray,
  FormBuilder,
  Validators,
  AbstractControl,
  FormGroupDirective,
} from '@angular/forms';
import { initFlowbite } from 'flowbite';
import { QuizQuestion } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-question-input',
  templateUrl: './question-input.component.html',
  styleUrls: ['./question-input.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class QuestionInputComponent {
  @Input() i!: any;
  @Input() question!: any;

  updateIsCorrectValue(optionIndex: number) {
    const optionsArray = this.question.get('options') as FormArray;
    optionsArray.controls.forEach((control, index) => {
      if (index !== optionIndex) {
        control?.get('isCorrect')?.setValue(false);
      }
    });
  }
}
