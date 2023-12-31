import { HttpErrorResponse } from '@angular/common/http';
import { ErrorHandler, Injectable, NgZone, inject } from '@angular/core';
import { SnackbarService } from './snackbar.service';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root',
})
export class GlobalErrorHandler implements ErrorHandler {
  private snackbar = inject(SnackbarService);
  private zone = inject(NgZone);
  private router = inject(Router);

  handleError(error: Error | HttpErrorResponse): void {
    let errorMsg = '';
    if (error instanceof HttpErrorResponse) {
      console.log('Error from the server', error);
      // if (error.status === 401) {
      //   this.zone.run(() => this.router.navigateByUrl('/auth/admin-login'));
      // }
      errorMsg = error.error.error || error.statusText;
    } else {
      console.log('Error from the client', error);
      errorMsg = error.message;
    }
    this.zone.run(() => this.snackbar.showError(errorMsg));
  }
}
