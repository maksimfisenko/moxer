import { Flex, Heading, Separator, Text, VStack } from "@chakra-ui/react";
import type { Template } from "../types/types";
import GenerateDataForm from "./GenerateDataForm";
import TemplateContentPreview from "./TemplateContentPreview";
import TemplateInfo from "./TemplateInfo";

interface TemplatePreviewProps {
  selectedTempl: Template | null;
}

const TemplatePreview = ({ selectedTempl }: TemplatePreviewProps) => {
  return (
    <VStack flex={1} pl={4} align={"stretch"}>
      {selectedTempl ? (
        <>
          <Heading size={"xl"} mb={2.5}>
            {selectedTempl.name}
          </Heading>

          <Separator />
          <TemplateInfo selectedTempl={selectedTempl} />
          <Separator />
          <GenerateDataForm templID={selectedTempl.id} />
          <Separator />
          <TemplateContentPreview content={selectedTempl.content} />
        </>
      ) : (
        <Flex flex={1} align={"center"} justify={"center"} h="100%">
          <Text flex={1} fontSize={"xl"} textAlign={"center"} maxW={"80%"}>
            Choose the template from the left panel or create a new one
          </Text>
        </Flex>
      )}
    </VStack>
  );
};

export default TemplatePreview;
