import type { AxiosError } from "axios";

export const getApiErrorMessage = (error: AxiosError<any>) => {
  return error.response?.data?.message || "Something went wrong.";
};
