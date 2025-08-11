import { Separator, VStack } from "@chakra-ui/react";
import type { Template } from "../types/types";
import TemplatesList from "./TemplatesList";
import { useState } from "react";

import MyTemplatesContainer from "./MyTemplatesContainer";

interface TemplatesListContainerProps {
  template: Template | null;
  setTemplate: React.Dispatch<React.SetStateAction<Template | null>>;
}

const TemplatesListContainer = ({
  template,
  setTemplate,
}: TemplatesListContainerProps) => {
  const [open, setOpen] = useState(false);

  return (
    <VStack
      w={"20%"}
      pr={4}
      align={"stretch"}
      borderRight={"2px solid"}
      borderColor={"gray.200"}
    >
      <MyTemplatesContainer open={open} setOpen={setOpen} />
      <Separator />
      <TemplatesList selectedTempl={template} setSelectedTempl={setTemplate} />
    </VStack>
  );
};

export default TemplatesListContainer;
