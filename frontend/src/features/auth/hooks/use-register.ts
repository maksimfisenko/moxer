import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { RegisterRequest } from "../types/types";
import api from "@/config/api-client";
import type { AxiosErrorResponseData } from "@/types/types";

const useRegister = () => {
  return useMutation<void, AxiosError<AxiosErrorResponseData>, RegisterRequest>(
    {
      mutationKey: ["register"],
      mutationFn: (registerRequest: RegisterRequest) =>
        api.post("api/v1/auth/register", registerRequest),
    }
  );
};

export { useRegister };
