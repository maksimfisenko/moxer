import { useGetCurrentUser } from "@/hooks/use-get-current-user";
import {
  Flex,
  Heading,
  HStack,
  IconButton,
  Spinner,
  Tag,
} from "@chakra-ui/react";
import { useQueryClient } from "@tanstack/react-query";
import { LuCircleArrowRight, LuCircleUser } from "react-icons/lu";
import { useNavigate } from "react-router";

const Header = () => {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { data, isLoading, isError, error } = useGetCurrentUser();

  const handleLogout = () => {
    localStorage.removeItem("token");
    queryClient.removeQueries({ queryKey: ["me"] });
    navigate("/login");
  };

  if (isLoading) return <Spinner />;
  if (isError) return <div>{error.response?.data.message}</div>;
  if (!data) return null;

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
          <Tag.Label>{data.email}</Tag.Label>
        </Tag.Root>
        <IconButton
          aria-label="Log out"
          variant={"surface"}
          colorPalette={"red"}
          onClick={handleLogout}
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
