import { Component, inject } from '@angular/core';
import { initFlowbite } from 'flowbite';
import { TokenService } from './shared/data-access/token.service';
import { Tokens } from './shared/interfaces';
import { AuthService } from './shared/data-access/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent {
  private tokenservice = inject(TokenService);
  private authService = inject(AuthService);

  ngOnInit(): void {
    initFlowbite();

    const userToken = localStorage.getItem(Tokens.UserToken);

    if (userToken) {
      if (!this.tokenservice.isTokenExpired(userToken)) {
        this.authService.login();
      }
    }

    const adminToken = localStorage.getItem(Tokens.AdminToken);

    if (adminToken) {
      if (!this.tokenservice.isTokenExpired(adminToken)) {
        this.authService.adminLogin();
      }
    }
  }
}
