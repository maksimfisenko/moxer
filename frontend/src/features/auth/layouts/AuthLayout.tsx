import { Toaster } from "@/components/ui/toaster";
import { Flex } from "@chakra-ui/react";
import type React from "react";

interface AuthLayoutProps {
  children: React.ReactNode;
}

const AuthLayout = ({ children }: AuthLayoutProps) => {
  return (
    <>
      <Flex
        flex={1}
        height={"100%"}
        align={"center"}
        justify={"center"}
        backgroundColor={"blue.50"}
      >
        {children}
      </Flex>
      <Toaster />
    </>
  );
};

export default AuthLayout;
