import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Output,
  inject,
} from '@angular/core';
import {
  FormBuilder,
  Validators,
  AbstractControlOptions,
} from '@angular/forms';
import { CustomValidationService } from 'src/app/auth/ui/signup/custom-validation';
import { ISignup } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-upsert-input',
  templateUrl: './upsert-input.component.html',
  styleUrls: ['./upsert-input.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UpsertInputComponent {
  private fb = inject(FormBuilder);
  private customValidator = inject(CustomValidationService);

  @Output() onSignupSubmit: EventEmitter<ISignup> = new EventEmitter();

  signupForm = this.fb.group(
    {
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      phone: ['', [Validators.required]],
      password: [
        '',
        Validators.compose([
          Validators.required,
          this.customValidator.patternValidator(),
        ]),
      ],
      confirmPassword: ['', [Validators.required]],
    },
    {
      validator: this.customValidator.MatchPassword(
        'password',
        'confirmPassword'
      ),
    } as AbstractControlOptions
  );

  get signupFormControl() {
    return this.signupForm.controls;
  }

  onSubmit() {
    if(this.signupForm.valid){
      this.onSignupSubmit.emit(this.signupForm.value as ISignup);
    }
  }
}
