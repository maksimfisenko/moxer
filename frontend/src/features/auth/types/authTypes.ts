import type z from "zod";
import type {
  loginFormSchema,
  registerFormSchema,
} from "../schemas/authSchemas";

export type LoginFormData = z.infer<typeof loginFormSchema>;
export type RegisterFormData = z.infer<typeof registerFormSchema>;

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
}
