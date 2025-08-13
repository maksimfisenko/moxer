import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { RegisterRequest } from "../types/types";
import type { AxiosErrorResponseData } from "@/types/types";
import publicAPI from "@/config/publicAPI";

const useRegister = () => {
  return useMutation<void, AxiosError<AxiosErrorResponseData>, RegisterRequest>(
    {
      mutationKey: ["register"],
      mutationFn: (registerRequest: RegisterRequest) =>
        publicAPI.post("/auth/register", registerRequest),
    }
  );
};

export { useRegister };
