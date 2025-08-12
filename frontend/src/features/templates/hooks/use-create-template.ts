import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { CreateTemplateRequest } from "../types/types";
import api from "@/config/api-client";
import type { AxiosErrorResponseData } from "@/types/types";

const useCreateTemplate = () => {
  return useMutation<
    void,
    AxiosError<AxiosErrorResponseData>,
    CreateTemplateRequest
  >({
    mutationKey: ["create-template"],
    mutationFn: (req: CreateTemplateRequest) =>
      api.post("api/v1/private/templates", req)
  });
};

export { useCreateTemplate };
