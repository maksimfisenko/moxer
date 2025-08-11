import { Heading, HStack } from "@chakra-ui/react";
import AddTemplateButton from "./AddTemplateButton";
import AddTemplateDrawer from "./AddTemplateDrawer";

interface MyTemplatesContainerProps {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const MyTemplatesContainer = ({ open, setOpen }: MyTemplatesContainerProps) => {
  return (
    <HStack justify={"space-between"} align={"center"} mb={2}>
      <Heading size={"xl"}>My templates</Heading>
      <AddTemplateButton setOpen={setOpen} />
      <AddTemplateDrawer open={open} setOpen={setOpen} />
    </HStack>
  );
};

export default MyTemplatesContainer;
