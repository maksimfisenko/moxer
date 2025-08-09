import { fullErrorMessages } from "@/stores/stores";

export const getFullErrorMessage = (shortErrorMessage: string | undefined) => {
  return fullErrorMessages[shortErrorMessage!] || shortErrorMessage;
};
