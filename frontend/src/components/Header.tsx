import type { User } from "@/types/types";
import { Flex, Heading, HStack } from "@chakra-ui/react";
import LogoutButton from "./LogoutButton";
import UserInfo from "./UserInfo";

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
      <Heading size={"3xl"}>MOXER</Heading>
      <HStack>
        <UserInfo email={user.email} />
        <LogoutButton onButtonClick={onButtonClick} />
      </HStack>
    </Flex>
  );
};

export default Header;
