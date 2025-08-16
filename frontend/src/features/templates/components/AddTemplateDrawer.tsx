import { CloseButton, Drawer, Portal } from "@chakra-ui/react";
import { Toaster } from "@/components/ui/toaster";
import { useQueryClient } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import { useCreateTemplate } from "../hooks/useCreateTemplate";
import type { CreateTemplateRequest } from "../types/templatesTypes";
import AddTemplateForm from "../forms/AddTemplateForm";
import { showErrorToast, showSuccessToast } from "@/shared/lib/toaster";
import { getApiErrorMessage } from "@/shared/lib/apiErrors";

interface AddTemplateDrawerProps {
  isOpen: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const AddTemplateDrawer = ({ isOpen, setOpen }: AddTemplateDrawerProps) => {
  const { mutate, isPending } = useCreateTemplate();
  const queryClient = useQueryClient();

  const handleCreateTemplate = (req: CreateTemplateRequest) => {
    mutate(req, {
      onSuccess: () => {
        showSuccessToast("Successfully created a new template!");
        queryClient.invalidateQueries({ queryKey: ["get-templates"] });
        setOpen(false);
      },
      onError: (error: AxiosError<any>) => {
        showErrorToast(getApiErrorMessage(error));
      },
    });
  };

  return (
    <>
      <Drawer.Root
        open={isOpen}
        onOpenChange={(e) => setOpen(e.open)}
        size={"full"}
        placement={"start"}
      >
        <Portal>
          <Drawer.Backdrop />
          <Drawer.Positioner>
            <Drawer.Content bgColor={"blue.50"}>
              <Drawer.Body>
                <AddTemplateForm
                  onFormSubmit={handleCreateTemplate}
                  isLoading={isPending}
                />
              </Drawer.Body>

              <Drawer.CloseTrigger asChild>
                <CloseButton size="sm" />
              </Drawer.CloseTrigger>
            </Drawer.Content>
          </Drawer.Positioner>
        </Portal>
      </Drawer.Root>
      <Toaster />
    </>
  );
};

export default AddTemplateDrawer;
