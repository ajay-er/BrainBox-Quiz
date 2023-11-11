import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { take, map } from 'rxjs';
import { AuthService } from '../data-access/auth.service';

export const adminUnauthGuard: CanActivateFn = (route, state) => {
  const authService = inject(AuthService);
  const router = inject(Router);

  return authService.isAdmin$.pipe(
    take(1),
    map((isAdminLoggedIn) => {
      if (!isAdminLoggedIn) {
        return true;
      } else {
        router.navigateByUrl('/admin/users');
        return false;
      }
    })
  );
};
