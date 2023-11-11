import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { ILogin, ISignup } from 'src/app/shared/interfaces';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root',
})
export class AuthApiService {
  private baseUrl = environment.apiUrl;
  private http = inject(HttpClient);

  submitLogin(loginData: ILogin) {
    return this.http.post(`${this.baseUrl}/users/login`, loginData);
  }

  submitSignup(loginData: ISignup) {
    return this.http.post(`${this.baseUrl}/users/signup`, loginData);
  }
}
