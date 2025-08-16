import { Flex } from "@chakra-ui/react";
import type React from "react";

interface MainLayoutProps {
  children: React.ReactNode;
}

const MainLayout = ({ children }: MainLayoutProps) => {
  return (
    <Flex h={"100%"} direction={"column"} bgColor={"blue.50"}>
      {children}
    </Flex>
  );
};

export default MainLayout;
