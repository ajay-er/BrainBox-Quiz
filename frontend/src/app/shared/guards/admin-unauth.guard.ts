import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { take, map } from 'rxjs';
import { AuthService } from '../data-access/auth.service';

export const adminUnauthGuard: CanActivateFn = (route, state) => {
  const authService = inject(AuthService);

  return authService.isAdmin$.pipe(
    take(1),
    map((isAdminLoggedIn) => {
      if (!isAdminLoggedIn) {
        return true;
      } else {
        return false;
      }
    })
  );
};
