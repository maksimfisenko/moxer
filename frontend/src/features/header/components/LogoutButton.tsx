import { IconButton } from "@chakra-ui/react";
import { LuCircleArrowRight } from "react-icons/lu";

interface LogoutButtonProps {
  onButtonClick: () => void;
}

const LogoutButton = ({ onButtonClick }: LogoutButtonProps) => {
  return (
    <IconButton
      aria-label="Log out"
      variant={"surface"}
      colorPalette={"red"}
      onClick={onButtonClick}
      size={"xs"}
      rounded={"lg"}
    >
      <LuCircleArrowRight />
    </IconButton>
  );
};

export default LogoutButton;
