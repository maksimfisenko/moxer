import { IconButton } from "@chakra-ui/react";
import { LuCirclePlus } from "react-icons/lu";

interface AddTemplateButtonProps {
  setDrawerOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const AddTemplateButton = ({ setDrawerOpen }: AddTemplateButtonProps) => {
  return (
    <IconButton
      size={"md"}
      colorPalette={"green"}
      rounded={"md"}
      variant={"surface"}
      onClick={() => setDrawerOpen(true)}
    >
      Add a new template
      <LuCirclePlus />
    </IconButton>
  );
};

export default AddTemplateButton;
