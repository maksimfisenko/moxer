import type { User } from "@/types/types";
import { Flex, Heading, HStack, IconButton, Tag } from "@chakra-ui/react";
import { LuCircleArrowRight, LuCircleUser } from "react-icons/lu";

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
        <Tag.Root size={"xl"} variant={"outline"} rounded={"lg"}>
          <Tag.StartElement>
            <LuCircleUser />
          </Tag.StartElement>
          <Tag.Label>{user.email}</Tag.Label>
        </Tag.Root>
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
      </HStack>
    </Flex>
  );
};

export default Header;
