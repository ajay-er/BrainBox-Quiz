export interface ILogin {
  email: string;
  password: string;
}
export interface ISignup {
  firstname: string;
  lastname: string;
  phone: string;
  email: string;
  password: string;
  confirmpassword: string;
}

export enum Tokens {
  AdminToken = 'adminToken',
  UserToken = 'userToken',
}

interface User {
  id: number;
  firstname: string;
  lastname: string;
  email: string;
  phone: string;
}

export interface UserAuthResponse {
  status_code: number;
  message: string;
  data: {
    Users: User;
    AccessToken: string;
    RefreshToken: string;
  };
  error: null | string;
}
