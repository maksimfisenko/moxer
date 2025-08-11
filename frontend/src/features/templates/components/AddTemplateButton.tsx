import { IconButton } from "@chakra-ui/react";
import { LuCirclePlus } from "react-icons/lu";

interface AddTemplateButtonProps {
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const AddTemplateButton = ({ setOpen }: AddTemplateButtonProps) => {
  return (
    <IconButton
      size={"xs"}
      colorPalette={"green"}
      rounded={"md"}
      variant={"surface"}
      onClick={() => setOpen(true)}
    >
      <LuCirclePlus />
    </IconButton>
  );
};

export default AddTemplateButton;
