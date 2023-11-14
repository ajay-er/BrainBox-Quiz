import { Injectable } from '@angular/core';
import { jwtDecode } from 'jwt-decode';

interface JwtPayload {
  role: string;
  exp: number;
}

@Injectable({
  providedIn: 'root',
})
export class TokenService {
  getUserRole(token: string): string {
    const decodedToken: JwtPayload = jwtDecode(token);
    return decodedToken.role;
  }

  private getDecodedToken(token: string) {
    const decodedToken: JwtPayload = jwtDecode(token);
    return decodedToken;
  }

  isTokenExpired(token: string): boolean {
    const decodedToken = this.getDecodedToken(token);
    const expirationTime = decodedToken.exp;
    const currentTimestamp = Math.floor(Date.now() / 1000);
    return expirationTime < currentTimestamp;
  }

  isAdmin(token: string): boolean {
    return this.getUserRole(token) === 'admin';
  }

  isUser(token: string): boolean {
    return this.getUserRole(token) === 'user';
  }
}
