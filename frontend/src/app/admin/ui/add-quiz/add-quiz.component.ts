import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Input,
  Output,
  inject,
} from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CreateQuiz, ICategory } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-add-quiz',
  templateUrl: './add-quiz.component.html',
  styleUrls: ['./add-quiz.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AddQuizComponent {
  @Input() categories: ICategory[] = [];
  @Output() submitAddQuizForm: EventEmitter<any> = new EventEmitter();

  private fb = inject(FormBuilder);

  addquizForm!: FormGroup;

  ngOnInit(): void {
    this.addquizForm = this.fb.group({
      quizName: ['', Validators.required],
      description: ['', Validators.required],
      categoryName: ['', Validators.required],
      questions: this.fb.array([this.createquiz()]),
    });
  }

  createquiz() {
    return this.fb.group({
      question: ['', Validators.required],
      options: this.fb.array(
        [
          this.createOption('', false),
          this.createOption('', false),
          this.createOption('', false),
          this.createOption('', false),
        ],
        Validators.required
      ),
    });
  }

  createOption(option: string, isCorrect: boolean) {
    return this.fb.group({
      option: [option, Validators.required],
      isCorrect: [isCorrect],
    });
  }

  get questions() {
    return (this.addquizForm.get('questions') as FormArray)
      .controls as FormGroup[];
  }

  addQuestion() {
    const questionsArray = this.addquizForm.get('questions') as FormArray;
    questionsArray.push(this.createquiz());
  }

  onSubmit() {
    if (this.addquizForm.valid) {
      this.submitAddQuizForm.emit(this.addquizForm.value as CreateQuiz);
    }
  }
}
