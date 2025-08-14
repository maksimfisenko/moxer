import { Flex } from "@chakra-ui/react";
import TemplatesListContainer from "./TemplatesListContainer";
import TemplatePreview from "./TemplatePreview";
import { useState } from "react";
import type { Template } from "../types/types";

const Content = () => {
  const [selectedTemplate, setSelectedTemplate] = useState<Template | null>(
    null
  );

  return (
    <Flex
      flex={1}
      p={4}
      bgColor={"white"}
      w={"80%"}
      alignSelf={"center"}
      shadow={"md"}
      rounded={"lg"}
      mt={2}
      mb={4}
    >
      <TemplatesListContainer
        selectedTemplate={selectedTemplate}
        setSelectedTemplate={setSelectedTemplate}
      />
      <TemplatePreview selectedTemplate={selectedTemplate} />
    </Flex>
  );
};

export default Content;
