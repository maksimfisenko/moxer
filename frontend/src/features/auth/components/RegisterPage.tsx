import RegisterForm from "./RegisterForm";
import type { RegisterRequest } from "../types/types";
import { useRegister } from "../hooks/use-register";
import { Toaster, toaster } from "@/components/ui/toaster";
import type { AxiosError } from "axios";
import { type AxiosErrorResponseData } from "@/types/types";
import { getFullErrorMessage } from "@/utils/utils";
import { useNavigate } from "react-router";

const RegisterPage = () => {
  const { mutate, isPending } = useRegister();
  const navigate = useNavigate();

  const handleRegister = (registerRequest: RegisterRequest) => {
    mutate(registerRequest, {
      onSuccess: () => {
        toaster.create({
          title: "Success",
          description: "Successfully registered new account!",
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

  const handleNavigateToLoginPage = () => {
    navigate("/login");
  };

  return (
    <>
      <RegisterForm
        isLoading={isPending}
        onFormSubmit={handleRegister}
        onLinkClick={handleNavigateToLoginPage}
      />
      <Toaster />
    </>
  );
};

export default RegisterPage;
