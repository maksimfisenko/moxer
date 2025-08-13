import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { LoginRequest, Token } from "../types/types";
import type { AxiosErrorResponseData } from "@/types/types";
import publicAPI from "@/config/publicAPI";

const useLogin = () => {
  return useMutation<Token, AxiosError<AxiosErrorResponseData>, LoginRequest>({
    mutationKey: ["login"],
    mutationFn: (loginRequest: LoginRequest) =>
      publicAPI
        .post<Token>("/auth/login", loginRequest)
        .then((response) => response.data),
  });
};

export { useLogin };
