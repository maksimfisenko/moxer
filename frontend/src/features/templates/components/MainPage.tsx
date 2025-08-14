import { Center, Flex, Spinner } from "@chakra-ui/react";
import Header from "@/components/Header";
import Content from "./Content";
import { useGetCurrentUser } from "@/hooks/use-get-current-user";
import { useNavigate } from "react-router";
import { useQueryClient } from "@tanstack/react-query";

const MainPage = () => {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const {
    data: user,
    isLoading: userIsLoading,
    isError: userIsError,
  } = useGetCurrentUser();

  const handleLogout = () => {
    localStorage.removeItem("token");
    queryClient.removeQueries({ queryKey: ["me"] });
    navigate("/login");
  };

  if (userIsLoading)
    return (
      <Center w="100%">
        <Spinner size={"xl"} />
      </Center>
    );

  if (userIsError || !user) {
    handleLogout();
    return null;
  }

  return (
    <Flex h={"100%"} direction={"column"} bgColor={"gray.100"}>
      <Header user={user} onButtonClick={handleLogout} />
      <Content />
    </Flex>
  );
};

export default MainPage;
