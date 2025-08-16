import { Separator, VStack } from "@chakra-ui/react";
import type { Template } from "../../types/templatesTypes";
import { useState } from "react";
import TemplatesHeader from "../TemplatesHeader";
import AddTemplatePanel from "./AddTemplatePanel";
import TemplatesList from "../TemplatesList";

interface TemplatesPanelProps {
  selectedTemplate: Template | null;
  setSelectedTemplate: React.Dispatch<React.SetStateAction<Template | null>>;
}

const TemplatesPanel = ({
  selectedTemplate,
  setSelectedTemplate,
}: TemplatesPanelProps) => {
  const [drawerIsOpen, setDrawerOpen] = useState(false);

  return (
    <VStack
      w={"20%"}
      pr={4}
      align={"stretch"}
      borderRight={"2px solid"}
      borderColor={"gray.200"}
    >
      <TemplatesHeader />
      <Separator />
      <AddTemplatePanel
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

export default TemplatesPanel;
