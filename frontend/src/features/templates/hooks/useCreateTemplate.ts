import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { CreateTemplateRequest } from "../types/templatesTypes";
import privateAPI from "@/shared/api/privateAPI";

const useCreateTemplate = () => {
  return useMutation<void, AxiosError<any>, CreateTemplateRequest>({
    mutationKey: ["create-template"],
    mutationFn: (req: CreateTemplateRequest) =>
      privateAPI.post("/templates", req),
  });
};

export { useCreateTemplate };
