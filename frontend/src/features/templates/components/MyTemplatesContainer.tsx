import { Heading, HStack } from "@chakra-ui/react";
import AddTemplateButton from "./AddTemplateButton";
import AddTemplateDrawer from "./AddTemplateDrawer";

interface MyTemplatesContainerProps {
  drawerIsOpen: boolean;
  setDrawerOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const MyTemplatesContainer = ({
  drawerIsOpen,
  setDrawerOpen,
}: MyTemplatesContainerProps) => {
  return (
    <HStack justify={"space-between"} align={"center"} mb={2}>
      <Heading size={"xl"}>My templates</Heading>
      <AddTemplateButton setDrawerOpen={setDrawerOpen} />
      <AddTemplateDrawer isOpen={drawerIsOpen} setOpen={setDrawerOpen} />
    </HStack>
  );
};

export default MyTemplatesContainer;
