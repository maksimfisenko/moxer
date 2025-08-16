import { Heading, HStack, Separator } from "@chakra-ui/react";
import Socials from "./Socials";

const AppInfo = () => {
  return (
    <HStack>
      <Heading size={"3xl"}>MOXER</Heading>
      <Separator orientation={"vertical"} height={6} />
      <Socials />
    </HStack>
  );
};

export default AppInfo;
