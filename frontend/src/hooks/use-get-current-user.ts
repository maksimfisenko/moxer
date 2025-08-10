import { useQuery } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { User } from "../types/types";
import api from "@/config/api-client";
import type { AxiosErrorResponseData } from "@/types/types";

const useGetCurrentUser = () => {
  return useQuery<User, AxiosError<AxiosErrorResponseData>>({
    queryKey: ["me"],
    queryFn: async () => (await api.get<User>("api/v1/auth/me")).data,
  });
};

export { useGetCurrentUser };
