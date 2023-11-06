import { Component, inject } from '@angular/core';
import { AuthService } from 'src/app/shared/data-access/auth.service';
import { SnackbarService } from 'src/app/shared/data-access/snackbar.service';
import { ILogin, ISignup } from 'src/app/shared/interfaces';

@Component({
  selector: 'app-auth-container',
  templateUrl: './auth-container.component.html',
  styleUrls: ['./auth-container.component.css'],
})
export class AuthContainerComponent {
  private authSevice = inject(AuthService);
  private snackBar = inject(SnackbarService);

  loginSubmit(data: ILogin) {
    this.authSevice.submitLogin(data).subscribe((res) => {
      console.log(res);
      this.snackBar.showSuccess('Login successfull');
    });
  }

  signupSubmit(data: ISignup) {
    this.authSevice.submitSignup(data).subscribe((res) => {
      console.log(res);
      this.snackBar.showSuccess('Signup successfull');
    });
  }
}
