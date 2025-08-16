import { useQuery } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import privateAPI from "@/shared/api/privateAPI";
import type { User } from "../types/headerTypes";

const useGetCurrentUser = () => {
  return useQuery<User, AxiosError<any>>({
    queryKey: ["me"],
    queryFn: async () => (await privateAPI.get<User>("/auth/me")).data,
    retry: 0,
  });
};

export { useGetCurrentUser };
