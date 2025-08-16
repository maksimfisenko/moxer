import { Center, Spinner } from "@chakra-ui/react";

const LoadingSpinner = () => {
  return (
    <Center w="100%" h="100%">
      <Spinner size={"xl"} />
    </Center>
  );
};

export default LoadingSpinner;
