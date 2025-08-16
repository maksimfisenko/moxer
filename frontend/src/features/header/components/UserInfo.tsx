import { Tag } from "@chakra-ui/react";
import { LuCircleUser } from "react-icons/lu";

interface UserInfoProps {
  email: string;
}

const UserInfo = ({ email }: UserInfoProps) => {
  return (
    <Tag.Root
      size={"xl"}
      variant={"surface"}
      rounded={"lg"}
      colorPalette={"blue"}
    >
      <Tag.StartElement>
        <LuCircleUser />
      </Tag.StartElement>
      <Tag.Label>{email}</Tag.Label>
    </Tag.Root>
  );
};

export default UserInfo;
