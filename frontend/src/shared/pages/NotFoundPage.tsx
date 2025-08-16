import { useNavigate } from "react-router";

import { Button, Flex, Heading, Text } from "@chakra-ui/react";

const NotFoundPage = () => {
  const navigate = useNavigate();

  return (
    <Flex
      direction={"column"}
      align={"center"}
      justify={"center"}
      height={"100%"}
      textAlign={"center"}
      gap={6}
      bgColor={"blue.50"}
    >
      <Heading fontSize={"6xl"} color={"blue.500"}>
        404
      </Heading>
      <Text fontSize={"xl"}>
        Oops! The page you are looking for does not exist.
      </Text>
      <Button colorPalette={"blue"} onClick={() => navigate("/")}>
        Go Home
      </Button>
    </Flex>
  );
};

export default NotFoundPage;
