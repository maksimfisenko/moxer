import z from "zod";

export const loginFormSchema = z.object({
  email: z.email("Invalid email address"),
  password: z.string().min(8, "Password should be at least 8 characters long"),
});

export const registerFormSchema = z.object({
  email: z.email("Invalid email address"),
  password: z.string().min(8, "Password should be at least 8 characters long"),
});
