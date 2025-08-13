import { useMutation } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { GenerateDataRequest, GeneratedData } from "../types/types";
import type { AxiosErrorResponseData } from "@/types/types";
import privateAPI from "@/config/privateAPI";

const useGenerateData = () => {
  return useMutation<
    GeneratedData,
    AxiosError<AxiosErrorResponseData>,
    { id: string; req: GenerateDataRequest }
  >({
    mutationKey: ["generate-data"],
    mutationFn: async ({ id, req }) => {
      const response = await privateAPI.post<GeneratedData>(
        `/templates/${id}/generate`,
        req
      );
      return response.data;
    },
  });
};

export { useGenerateData };
