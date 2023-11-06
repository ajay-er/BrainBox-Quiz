import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Output,
  inject,
} from '@angular/core';
import {
  AbstractControlOptions,
  FormBuilder,
  Validators,
} from '@angular/forms';
import { ISignup } from 'src/app/shared/interfaces';
import { CustomValidationService } from './custom-validation';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SignupComponent {
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

  onSubmit() {}
}
