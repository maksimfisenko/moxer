export interface User {
  id: string;
  email: string;
  created_at: Date;
  updated_at: Date;
}

export interface AxiosErrorResponseData {
  message: string;
}
