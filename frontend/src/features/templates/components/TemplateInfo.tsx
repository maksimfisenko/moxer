import { Heading, HStack } from "@chakra-ui/react";
import type { Template } from "../types/templatesTypes";

interface TemplateInfoProps {
  selectedTemplate: Template;
}

const TemplateInfo = ({ selectedTemplate }: TemplateInfoProps) => {
  return (
    <HStack justify={"space-between"}>
      <Heading size={"xl"}>{selectedTemplate.name}</Heading>
    </HStack>
  );
};

export default TemplateInfo;
