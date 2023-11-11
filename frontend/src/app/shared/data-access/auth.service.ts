import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Tokens } from '../interfaces';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private isAuthenticatedSubject = new BehaviorSubject<boolean>(false);
  isLoggedIn$ = this.isAuthenticatedSubject.asObservable();

  private isAdminSubject = new BehaviorSubject<boolean>(false);
  isAdmin$ = this.isAdminSubject.asObservable();

  login() {
    this.isAuthenticatedSubject.next(true);
  }

  logout() {
    window.localStorage.removeItem(Tokens.UserToken);
    this.isAuthenticatedSubject.next(false);
  }

  adminLogin() {
    this.isAdminSubject.next(true);
  }

  adminLogout() {
    window.localStorage.removeItem(Tokens.AdminToken);
    this.isAdminSubject.next(false);
  }
}
