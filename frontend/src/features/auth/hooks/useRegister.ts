import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import publicAPI from "@/shared/api/publicAPI";
import type { RegisterRequest } from "../types/authTypes";

const useRegister = () => {
  return useMutation<void, AxiosError<any>, RegisterRequest>({
    mutationKey: ["register"],
    mutationFn: (registerRequest: RegisterRequest) =>
      publicAPI.post("/auth/register", registerRequest),
  });
};

export { useRegister };
