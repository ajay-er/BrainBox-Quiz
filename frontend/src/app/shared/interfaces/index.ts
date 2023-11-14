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

interface Admin {
  id: number;
  firstname: string;
  lastname: string;
  email: string;
}

export interface AdminAuthResponse {
  status_code: number;
  message: string;
  data: {
    Admin: Admin;
    Token: string;
  };
  error: null | string;
}

export enum PageLayout {
  Admin = 'admin',
  User = 'user',
}

export interface ICategory {
  category_name: string;
  icon_svg: string;
  total_quizes?:number;
  id?:string
}

export interface HomePageResponse {
    Categories:ICategory[];
}

export interface CreateQuiz {
  categoryName: string;
  quizName: string;
  questions: QuizQuestion[];
}

export interface QuizQuestion {
  question: string;
  options: OptionValues[];
}

export interface OptionValues {
  option: string;
  isCorrect: boolean;
}
