import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { ILogin, ISignup } from '../interfaces';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private http = inject(HttpClient);
  private baseUrl = environment.apiUrl;

  submitLogin(loginData: ILogin) {
    return this.http.post(`${this.baseUrl}/user/login`, loginData);
  }
  submitSignup(loginData: ISignup) {
    return this.http.post(`${this.baseUrl}/user/signup`, loginData);
  }
}
