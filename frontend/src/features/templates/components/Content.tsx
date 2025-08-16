import { Flex } from "@chakra-ui/react";
import { useState } from "react";
import type { Template } from "../types/templatesTypes";
import TemplatesPanel from "./panels/TemplatesPanel";
import TemplatePanel from "./panels/TemplatePanel";

const Content = () => {
  const [selectedTemplate, setSelectedTemplate] = useState<Template | null>(
    null
  );

  return (
    <Flex
      flex={1}
      h={"100%"}
      p={4}
      bgColor={"white"}
      w={"80%"}
      alignSelf={"center"}
      shadow={"md"}
      rounded={"lg"}
      mt={2}
      mb={4}
    >
      <TemplatesPanel
        selectedTemplate={selectedTemplate}
        setSelectedTemplate={setSelectedTemplate}
      />
      <TemplatePanel selectedTemplate={selectedTemplate} />
    </Flex>
  );
};

export default Content;
