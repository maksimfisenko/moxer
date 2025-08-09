import type { LoginRequest } from "../types/types";
import { Toaster, toaster } from "@/components/ui/toaster";
import type { AxiosError } from "axios";
import { type AxiosErrorResponseData } from "@/types/types";
import { getFullErrorMessage } from "@/utils/utils";
import { useLogin } from "../hooks/use-login";
import LoginForm from "./LoginForm";

const LoginPage = () => {
  const { mutate, isPending } = useLogin();

  const handleLogin = (loginRequest: LoginRequest) => {
    mutate(loginRequest, {
      onSuccess: (token) => {
        console.log(token.token);
        toaster.create({
          title: "Success",
          description: "Successfully signed in account!",
          type: "success",
        });
      },
      onError: (error: AxiosError<AxiosErrorResponseData>) => {
        toaster.create({
          title: "Error",
          description: getFullErrorMessage(error.response?.data?.message),
          type: "error",
        });
      },
    });
  };

  return (
    <>
      <LoginForm isLoading={isPending} onFormSubmit={handleLogin} />
      <Toaster />
    </>
  );
};

export default LoginPage;
