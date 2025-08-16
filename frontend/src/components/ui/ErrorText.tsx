import { Center, Code } from "@chakra-ui/react";

interface ErrorTextProps {
  message: string;
  details: string | undefined;
}

const ErrorText = ({ message, details }: ErrorTextProps) => {
  return (
    <Center w="100%" h="100%">
      <Code
        size={"sm"}
        textAlign={"center"}
        colorPalette={"red"}
        variant={"surface"}
      >
        {details
          ? `${message} [${details}]. Please, try again.`
          : `${message}. Please, try again.`}
      </Code>
    </Center>
  );
};

export default ErrorText;
