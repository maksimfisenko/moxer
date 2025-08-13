import { useQuery } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { User } from "../types/types";
import type { AxiosErrorResponseData } from "@/types/types";
import privateAPI from "@/config/privateAPI";

const useGetCurrentUser = () => {
  return useQuery<User, AxiosError<AxiosErrorResponseData>>({
    queryKey: ["me"],
    queryFn: async () => (await privateAPI.get<User>("/auth/me")).data,
    retry: 0,
  });
};

export { useGetCurrentUser };
