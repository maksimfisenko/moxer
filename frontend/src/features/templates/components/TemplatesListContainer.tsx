import { Separator, VStack } from "@chakra-ui/react";
import type { Template } from "../types/types";
import TemplatesList from "./TemplatesList";
import { useState } from "react";

import MyTemplatesContainer from "./MyTemplatesContainer";

interface TemplatesListContainerProps {
  selectedTemplate: Template | null;
  setSelectedTemplate: React.Dispatch<React.SetStateAction<Template | null>>;
}

const TemplatesListContainer = ({
  selectedTemplate,
  setSelectedTemplate,
}: TemplatesListContainerProps) => {
  const [drawerIsOpen, setDrawerOpen] = useState(false);

  return (
    <VStack
      w={"20%"}
      pr={4}
      align={"stretch"}
      borderRight={"2px solid"}
      borderColor={"gray.200"}
    >
      <MyTemplatesContainer
        drawerIsOpen={drawerIsOpen}
        setDrawerOpen={setDrawerOpen}
      />
      <Separator />
      <TemplatesList
        selectedTemplate={selectedTemplate}
        setSelectedTemplate={setSelectedTemplate}
      />
    </VStack>
  );
};

export default TemplatesListContainer;
