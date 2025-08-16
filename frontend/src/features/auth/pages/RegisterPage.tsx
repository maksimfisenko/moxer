import type { AxiosError } from "axios";
import { useNavigate } from "react-router";
import { showErrorToast, showSuccessToast } from "@/shared/lib/toaster";
import { getApiErrorMessage } from "@/shared/lib/apiErrors";
import AuthLayout from "../layouts/AuthLayout";
import { useRegister } from "../hooks/useRegister";
import RegisterForm from "../forms/RegisterForm";
import type { RegisterRequest } from "../types/authTypes";

const RegisterPage = () => {
  const { mutate, isPending } = useRegister();
  const navigate = useNavigate();

  const handleRegister = (registerRequest: RegisterRequest) => {
    mutate(registerRequest, {
      onSuccess: () => {
        showSuccessToast("Successfully registered a new account!");
      },
      onError: (error: AxiosError<any>) => {
        showErrorToast(getApiErrorMessage(error));
      },
    });
  };

  const handleNavigateToLoginPage = () => {
    navigate("/login");
  };

  return (
    <AuthLayout>
      <RegisterForm
        isLoading={isPending}
        onFormSubmit={handleRegister}
        onLinkClick={handleNavigateToLoginPage}
      />
    </AuthLayout>
  );
};

export default RegisterPage;
