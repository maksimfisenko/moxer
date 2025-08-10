import {
  Heading,
  HStack,
  IconButton,
  Separator,
  VStack,
} from "@chakra-ui/react";
import type { Template } from "../types/types";
import { LuCirclePlus } from "react-icons/lu";
import TemplatesList from "./TemplatesList";

interface TemplatesListContainerProps {
  selectedTempl: Template | null;
  setSelectedTempl: React.Dispatch<React.SetStateAction<Template | null>>;
}

const TemplatesListContainer = ({
  selectedTempl,
  setSelectedTempl,
}: TemplatesListContainerProps) => {
  return (
    <VStack
      w={"20%"}
      pr={4}
      align={"stretch"}
      borderRight={"2px solid"}
      borderColor={"gray.200"}
    >
      <HStack justify={"space-between"} align={"center"} mb={2}>
        <Heading size={"xl"}>My templates</Heading>
        <IconButton
          size={"xs"}
          colorPalette={"green"}
          rounded={"md"}
          variant={"surface"}
        >
          <LuCirclePlus />
        </IconButton>
      </HStack>

      <Separator />

      <TemplatesList
        selectedTempl={selectedTempl}
        setSelectedTempl={setSelectedTempl}
      />
    </VStack>
  );
};

export default TemplatesListContainer;
