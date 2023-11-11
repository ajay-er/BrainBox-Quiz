import { Component, inject } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { filter } from 'rxjs';
import { SnackbarService } from 'src/app/shared/data-access/snackbar.service';
import { ILogin, ISignup } from 'src/app/shared/interfaces';
import { AuthApiService } from '../../data-access/auth-api.service';

@Component({
  selector: 'app-auth-container',
  templateUrl: './auth-container.component.html',
  styleUrls: ['./auth-container.component.css'],
})
export class AuthContainerComponent {
  private router = inject(Router);
  private authSevice = inject(AuthApiService);
  private snackBar = inject(SnackbarService);

  isLoginRoute = false;
  isSignupRoute = false;

  constructor() {
    this.router.events
      .pipe(filter((event) => event instanceof NavigationEnd))
      .subscribe((event: any) => {
        this.isLoginRoute = event.url === '/auth/login';
        this.isSignupRoute = event.url === '/auth/signup';
      });
  }

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
