import { Component, inject } from '@angular/core';
import { AuthService } from 'src/app/shared/data-access/auth.service';
import { SnackbarService } from 'src/app/shared/data-access/snackbar.service';
import {
  ILogin,
  UserAuthResponse,
  Tokens,
  AdminAuthResponse,
} from 'src/app/shared/interfaces';
import { AdminService } from '../../data-access/admin.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-admin-login',
  templateUrl: './admin-login.component.html',
  styleUrls: ['./admin-login.component.css'],
})
export class AdminLoginComponent {
  private auth = inject(AuthService);
  private snackBar = inject(SnackbarService);
  private adminService = inject(AdminService);
  private router = inject(Router);

  loginSubmit(data: ILogin) {
    this.adminService.submitLogin(data).subscribe((res: AdminAuthResponse) => {
      console.log(res);
      localStorage.setItem(Tokens.AdminToken, res.data.Token);
      this.auth.login();
      this.router.navigateByUrl('/admin/users')
      this.snackBar.showSuccess('Admin Login successfull');
    });
  }
}
