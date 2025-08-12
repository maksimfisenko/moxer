import type { User } from "@/types/types";
import { Flex, Heading, HStack, Link, Separator } from "@chakra-ui/react";
import LogoutButton from "./LogoutButton";
import UserInfo from "./UserInfo";
import { LuGithub } from "react-icons/lu";
import { SocialIcon } from "react-social-icons";

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
      <HStack>
        <Heading size={"3xl"}>MOXER</Heading>
        <Separator orientation={"vertical"} height={6} />
        <Link href="https://github.com/maksimfisenko/moxer" target="_blank">
          <LuGithub />
        </Link>
        <Link>
          <SocialIcon
            target="_blank"
            url="https://t.me/maximfisenko"
            style={{ height: "20px", width: "20px" }}
          />
        </Link>
      </HStack>

      <HStack>
        <UserInfo email={user.email} />
        <LogoutButton onButtonClick={onButtonClick} />
      </HStack>
    </Flex>
  );
};

export default Header;
