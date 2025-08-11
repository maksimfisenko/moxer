import { CloseButton, Drawer, Portal } from "@chakra-ui/react";
import AddTemplateForm from "./AddTemplateForm";
import { useCreateTemplate } from "../hooks/use-create-template";
import type { CreateTemplateRequest } from "../types/types";
import { toaster, Toaster } from "@/components/ui/toaster";
import { useQueryClient } from "@tanstack/react-query";
import type { AxiosError } from "axios";
import type { AxiosErrorResponseData } from "@/types/types";
import { getFullErrorMessage } from "@/utils/utils";

interface AddTemplateDrawerProps {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const AddTemplateDrawer = ({ open, setOpen }: AddTemplateDrawerProps) => {
  const { mutate, isPending } = useCreateTemplate();
  const queryClient = useQueryClient();

  const handleCreateTemplate = (req: CreateTemplateRequest) => {
    mutate(req, {
      onSuccess: () => {
        toaster.create({
          title: "Success",
          description: "Successfully created a new template!",
          type: "success",
        });
        queryClient.invalidateQueries({ queryKey: ["get-templates"] });
        setOpen(false);
      },
      onError: (error: AxiosError<AxiosErrorResponseData>) => {
        toaster.create({
          title: "Error",
          description: getFullErrorMessage(error.response?.data?.message),
          type: "error",
        });
      },
    });
  };

  return (
    <>
      <Drawer.Root
        open={open}
        onOpenChange={(e) => setOpen(e.open)}
        size={"full"}
        placement={"start"}
      >
        <Portal>
          <Drawer.Backdrop />
          <Drawer.Positioner>
            <Drawer.Content>
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
