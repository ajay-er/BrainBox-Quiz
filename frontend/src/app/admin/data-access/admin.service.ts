import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';
import { AdminAuthResponse, ILogin } from 'src/app/shared/interfaces';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root',
})
export class AdminService {
  private baseUrl = environment.apiUrl;
  private http = inject(HttpClient);

  submitLogin(loginData: ILogin): Observable<AdminAuthResponse> {
    return this.http.post<AdminAuthResponse>(
      `${this.baseUrl}/admin/login`,
      loginData
    );
  }
}
