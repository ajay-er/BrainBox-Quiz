import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';
import { ILogin, ISignup, UserAuthResponse } from 'src/app/shared/interfaces';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root',
})
export class AuthApiService {
  private baseUrl = environment.apiUrl;
  private http = inject(HttpClient);

  submitLogin(loginData: ILogin): Observable<UserAuthResponse> {
    return this.http.post<UserAuthResponse>(
      `${this.baseUrl}/users/login`,
      loginData
    );
  }

  submitSignup(loginData: ISignup): Observable<UserAuthResponse> {
    return this.http.post<UserAuthResponse>(
      `${this.baseUrl}/users/signup`,
      loginData
    );
  }
}
