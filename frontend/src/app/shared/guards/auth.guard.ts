import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../data-access/auth.service';
import { map, take } from 'rxjs';

export const authGuard: CanActivateFn = (route, state) => {
  const authService = inject(AuthService);
  const router = inject(Router);

  return authService.isLoggedIn$.pipe(
    take(1),
    map((isUserLoggedIn) => {
      if (isUserLoggedIn) {
        return true;
      } else {
        router.navigateByUrl('/auth/login');
        return false;
      }
    })
  );
};
