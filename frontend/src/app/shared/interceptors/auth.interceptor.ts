import { HttpInterceptorFn, HttpRequest } from '@angular/common/http';
import { Tokens } from '../interfaces';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const useAdminToken = shouldUseAdminToken(req);
  const token = useAdminToken ? getAdminToken() : getUserToken();

  if (token) {
    const modifiedRequest = req.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`,
      },
    });
    return next(modifiedRequest);
  }

  return next(req);
};

function shouldUseAdminToken(req: HttpRequest<unknown>): boolean {
  return req.url.startsWith('/admin');
}

function getAdminToken(): string | null {
  return localStorage.getItem(Tokens.AdminToken);
}

function getUserToken(): string | null {
  return localStorage.getItem(Tokens.UserToken);
}
