import { Flex } from "@chakra-ui/react";
import Header from "@/components/Header";
import Content from "./Content";

const MainPage = () => {
  return (
    <Flex flex={1} direction={"column"} bgColor={"gray.100"}>
      <Header />
      <Content />
    </Flex>
  );
};

export default MainPage;
