import { Center, Text } from "@chakra-ui/react";

interface AdvisoryTextProps {
  message: string;
}

const AdvisoryText = ({ message }: AdvisoryTextProps) => {
  return (
    <Center w="100%" h="100%">
      <Text textAlign={"center"} color={"blackAlpha.500"} textStyle={"lg"}>
        {message}
      </Text>
    </Center>
  );
};

export default AdvisoryText;
