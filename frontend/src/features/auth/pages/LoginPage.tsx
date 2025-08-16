import type { AxiosError } from "axios";
import { useNavigate } from "react-router";
import { showErrorToast } from "@/shared/lib/toaster";
import { getApiErrorMessage } from "@/shared/lib/apiErrors";
import { useLogin } from "../hooks/useLogin";
import type { LoginRequest } from "../types/authTypes";
import AuthLayout from "../layouts/AuthLayout";
import LoginForm from "../forms/LoginForm";

const LoginPage = () => {
  const { mutate, isPending } = useLogin();
  const navigate = useNavigate();

  const handleLogin = (loginRequest: LoginRequest) => {
    mutate(loginRequest, {
      onSuccess: (token) => {
        localStorage.setItem("token", token.token);
        navigate("/");
      },
      onError: (error: AxiosError<any>) => {
        showErrorToast(getApiErrorMessage(error));
      },
    });
  };

  const handleNavigateToRegisterPage = () => {
    navigate("/register");
  };

  return (
    <AuthLayout>
      <LoginForm
        isLoading={isPending}
        onFormSubmit={handleLogin}
        onLinkClick={handleNavigateToRegisterPage}
      />
    </AuthLayout>
  );
};

export default LoginPage;
