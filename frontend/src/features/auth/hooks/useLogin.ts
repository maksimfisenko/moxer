import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import publicAPI from "@/shared/api/publicAPI";
import type { LoginRequest } from "../types/authTypes";
import type { Token } from "@/shared/types/common";

const useLogin = () => {
  return useMutation<Token, AxiosError<any>, LoginRequest>({
    mutationKey: ["login"],
    mutationFn: (loginRequest: LoginRequest) =>
      publicAPI
        .post<Token>("/auth/login", loginRequest)
        .then((response) => response.data),
  });
};

export { useLogin };
