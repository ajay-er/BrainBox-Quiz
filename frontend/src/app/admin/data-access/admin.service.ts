import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { ICategory, ISignup } from 'src/app/shared/interfaces';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root',
})
export class AdminService {
  private baseUrl = environment.apiUrl;
  private http = inject(HttpClient);

  addNewCategory(data: ICategory) {    
    return this.http.post(`${this.baseUrl}/admin/category`, data);
  }

  addNewUser(user: ISignup) {
    return this.http.post(`${this.baseUrl}/users/signup`, user);
  }

  updateUserData(user: ISignup) {
    return this.http.post(`${this.baseUrl}/users/update-user`, user);
  }

  getAllCategories(){
    return this.http.get(`${this.baseUrl}/users/`)
  }

  addNewQuiz(data:any){
    return this.http.post(`${this.baseUrl}/admin/quizes`,data)
  }
}
