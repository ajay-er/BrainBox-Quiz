import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { take, map } from 'rxjs';
import { AuthService } from '../data-access/auth.service';

export const unauthGuard: CanActivateFn = (route, state) => {
  const authService = inject(AuthService);
  return authService.isLoggedIn$.pipe(
    take(1),
    map((isUserLoggedIn) => {
      if (isUserLoggedIn) {
        return false;
      } else {
        return true;
      }
    })
  );
};
