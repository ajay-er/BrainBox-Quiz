import {
  ChangeDetectionStrategy,
  Component,
  EventEmitter,
  Output,
  inject,
} from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { ILogin } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginComponent {
  private fb = inject(FormBuilder);
  @Output() onLoginSubmit: EventEmitter<ILogin> = new EventEmitter();
  
  loginForm = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', Validators.required],
  });

  onSubmit() {
    const loginData: ILogin = {
      email: this.loginForm.get('email')?.value || '',
      password: this.loginForm.get('password')?.value || '',
    };
    this.onLoginSubmit.emit(loginData);
  }
}
