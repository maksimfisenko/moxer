import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { GenerateDataRequest, GeneratedData } from "../types/types";
import api from "@/config/api-client";
import type { AxiosErrorResponseData } from "@/types/types";

const useGenerateData = () => {
  return useMutation<
    GeneratedData,
    AxiosError<AxiosErrorResponseData>,
    { id: string; req: GenerateDataRequest }
  >({
    mutationKey: ["generate-data"],
    mutationFn: async ({ id, req }) => {
      const response = await api.post<GeneratedData>(
        `api/v1/templates/${id}/generate`,
        req
      );
      return response.data;
    },
  });
};

export { useGenerateData };
