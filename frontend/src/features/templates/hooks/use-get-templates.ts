import { useQuery } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { Template } from "../types/types";
import type { AxiosErrorResponseData } from "@/types/types";
import privateAPI from "@/config/privateAPI";

const useGetTemplates = () => {
  return useQuery<Template[], AxiosError<AxiosErrorResponseData>>({
    queryKey: ["get-templates"],
    queryFn: async () => {
      const response = await privateAPI.get<Template[]>("/templates");
      return response.data.map((t) => ({
        ...t,
        created_at: new Date(t.created_at),
        updated_at: new Date(t.updated_at),
      }));
    },
    retry: 0,
  });
};

export { useGetTemplates };
