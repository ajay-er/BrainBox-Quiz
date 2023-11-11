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
