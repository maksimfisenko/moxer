import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { CreateTemplateRequest } from "../types/types";
import type { AxiosErrorResponseData } from "@/types/types";
import privateAPI from "@/config/privateAPI";

const useCreateTemplate = () => {
  return useMutation<
    void,
    AxiosError<AxiosErrorResponseData>,
    CreateTemplateRequest
  >({
    mutationKey: ["create-template"],
    mutationFn: (req: CreateTemplateRequest) =>
      privateAPI.post("/templates", req),
  });
};

export { useCreateTemplate };
