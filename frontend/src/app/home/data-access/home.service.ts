import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root',
})
export class HomeService {
  private apiBaseUrl = environment.apiUrl;
  private http = inject(HttpClient);

  getAllCategories() {
    return this.http.get(`${this.apiBaseUrl}`);
  }
}
