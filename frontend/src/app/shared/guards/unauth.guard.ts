import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { take, map } from 'rxjs';
import { AuthService } from '../data-access/auth.service';

export const unauthGuard: CanActivateFn = (route, state) => {
  const authService = inject(AuthService);
  const router = inject(Router);
  return authService.isLoggedIn$.pipe(
    take(1),
    map((isUserLoggedIn) => {
      if (isUserLoggedIn) {
        router.navigateByUrl('/home')
        return false;
      } else {
        return true;
      }
    })
  );
};
