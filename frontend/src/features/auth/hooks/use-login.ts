import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { LoginRequest, Token } from "../types/types";
import api from "@/config/api-client";
import type { AxiosErrorResponseData } from "@/types/types";

const useLogin = () => {
  return useMutation<Token, AxiosError<AxiosErrorResponseData>, LoginRequest>({
    mutationKey: ["login"],
    mutationFn: (loginRequest: LoginRequest) =>
      api
        .post<Token>("api/v1/public/auth/login", loginRequest)
        .then((response) => response.data),
  });
};

export { useLogin };
