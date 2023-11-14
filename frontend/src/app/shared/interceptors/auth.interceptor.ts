import { Injectable } from '@angular/core';
import {
  HttpInterceptor,
  HttpHandler,
  HttpRequest,
} from '@angular/common/http';
import { Tokens } from '../interfaces';
import { Router } from '@angular/router';
import { catchError, throwError } from 'rxjs';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  constructor(private router: Router) {}

  intercept(req: HttpRequest<unknown>, next: HttpHandler) {
    const useAdminToken = this.shouldUseAdminToken();

    const token = useAdminToken ? this.getAdminToken() : this.getUserToken();

    if (token) {
      const modifiedRequest = req.clone({
        setHeaders: {
          Authorization: `Bearer ${token}`,
        },
      });

      return next.handle(modifiedRequest).pipe(
        catchError((error) => {
          if (error.status === 401) {
            console.log('evide work aayi chetta..');
            this.router.navigate(['/auth/login']);
          }
          return throwError(() => error);
        })
      );
    }

    return next.handle(req);
  }

  private shouldUseAdminToken(): boolean {
    const currentRoute = this.router.url;
    return currentRoute.startsWith('/admin');
  }

  private getAdminToken(): string | null {
    const token = window.localStorage.getItem(Tokens.AdminToken);
    return token;
  }

  private getUserToken(): string | null {
    const token = window.localStorage.getItem(Tokens.UserToken);
    return token;
  }
}
