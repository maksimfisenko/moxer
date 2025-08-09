import { useState } from "react";
import RegisterForm from "./RegisterForm";
import api from "@/config/api-client";
import type { RegisterRequest } from "../types/types";

const RegisterPage = () => {
  const [error, setError] = useState<boolean>(false);
  const [success, setSuccess] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);

  const handleRegister = (registerRequest: RegisterRequest) => {
    setError(false);
    setSuccess(false);
    setLoading(true);

    api
      .post("api/v1/auth/register", registerRequest)
      .then(() => setSuccess(true))
      .catch(() => setError(true))
      .finally(() => setLoading(false));
  };

  return (
    <RegisterForm
      isError={error}
      isSuccess={success}
      isLoading={loading}
      onFormSubmit={handleRegister}
    />
  );
};

export default RegisterPage;
