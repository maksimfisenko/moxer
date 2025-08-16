import { Badge, HStack, Separator } from "@chakra-ui/react";

interface SeparatorWithTextProps {
  text: string;
}

const SeparatorWithText = ({ text }: SeparatorWithTextProps) => {
  return (
    <HStack>
      <Separator flex="1" />
      <Badge flexShrink="0" size={"md"} colorPalette={"blue"}>
        {text}
      </Badge>
      <Separator flex="1" />
    </HStack>
  );
};

export default SeparatorWithText;
