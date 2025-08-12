import { useQuery } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { Template } from "../types/types";
import api from "@/config/api-client";
import type { AxiosErrorResponseData } from "@/types/types";

const useGetTemplates = () => {
  return useQuery<Template[], AxiosError<AxiosErrorResponseData>>({
    queryKey: ["get-templates"],
    queryFn: async () => {
      const response = await api.get<Template[]>("api/v1/private/templates");
      return response.data.map((t) => ({
        ...t,
        created_at: new Date(t.created_at),
        updated_at: new Date(t.updated_at),
      }));
    },
  });
};

export { useGetTemplates };
