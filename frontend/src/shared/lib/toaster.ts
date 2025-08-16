import { toaster } from "@/components/ui/toaster";

export const showErrorToast = (description: string) => {
  toaster.create({
    title: "Error",
    description: description,
    type: "error",
  });
};

export const showSuccessToast = (description: string) => {
  toaster.create({
    title: "Success",
    description: description,
    type: "success",
  });
};
