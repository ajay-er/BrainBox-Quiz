import { Component, inject } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { filter } from 'rxjs';
import { SnackbarService } from 'src/app/shared/data-access/snackbar.service';
import {
  AdminAuthResponse,
  ILogin,
  ISignup,
  Tokens,
  UserAuthResponse,
} from 'src/app/shared/interfaces';
import { AuthApiService } from '../../data-access/auth-api.service';
import { AuthService } from 'src/app/shared/data-access/auth.service';

@Component({
  selector: 'app-auth-container',
  templateUrl: './auth-container.component.html',
  styleUrls: ['./auth-container.component.css'],
})
export class AuthContainerComponent {
  private router = inject(Router);
  private authApi = inject(AuthApiService);
  private auth = inject(AuthService);
  private snackBar = inject(SnackbarService);

  isLoginRoute = false;
  isSignupRoute = false;
  isAdminLoginRoute = false;

  constructor() {
    this.router.events
      .pipe(filter((event) => event instanceof NavigationEnd))
      .subscribe((event: any) => {
        this.isLoginRoute = event.url === '/auth/login';
        this.isSignupRoute = event.url === '/auth/signup';
        this.isAdminLoginRoute = event.url === '/auth/admin-login';
      });
  }

  loginSubmit(data: ILogin) {
    this.authApi.submitLogin(data).subscribe((res: UserAuthResponse) => {
      console.log(res);
      localStorage.setItem(Tokens.UserToken, res.data.AccessToken);
      this.auth.login();
      this.router.navigateByUrl('/home');
      this.snackBar.showSuccess('Login successfull');
    });
  }

  signupSubmit(data: ISignup) {
    this.authApi.submitSignup(data).subscribe((res: UserAuthResponse) => {
      console.log(res);
      localStorage.setItem(Tokens.UserToken, res.data.AccessToken);
      this.auth.login();
      this.router.navigateByUrl('/home');
      this.snackBar.showSuccess('Signup successfull');
    });
  }

  adminLoginSubmit(data: ILogin) {
    this.authApi.submitAdminLogin(data).subscribe((res: AdminAuthResponse) => {
      console.log(res);
      localStorage.setItem(Tokens.AdminToken, res.data.Token);
      this.auth.adminLogin();
      this.router.navigateByUrl('/admin/users');
      this.snackBar.showSuccess('Admin Login successfull');
    });
  }
}
