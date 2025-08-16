import { Flex, HStack } from "@chakra-ui/react";
import type { User } from "../types/headerTypes";
import AppInfo from "./AppInfo";
import UserInfo from "./UserInfo";
import LogoutButton from "./LogoutButton";

interface HeaderProps {
  user: User;
  onButtonClick: () => void;
}

const Header = ({ user, onButtonClick }: HeaderProps) => {
  return (
    <Flex
      p={4}
      as={"header"}
      justify={"space-between"}
      align={"center"}
      bgColor={"white"}
      w={"80%"}
      alignSelf={"center"}
      shadow={"md"}
      rounded={"lg"}
      mt={4}
      mb={2}
    >
      <AppInfo />
      <HStack>
        <UserInfo email={user.email} />
        <LogoutButton onButtonClick={onButtonClick} />
      </HStack>
    </Flex>
  );
};

export default Header;
