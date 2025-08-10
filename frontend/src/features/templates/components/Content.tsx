import { Flex } from "@chakra-ui/react";
import TemplatesListContainer from "./TemplatesListContainer";
import TemplatePreview from "./TemplatePreview";
import { useState } from "react";
import type { Template } from "../types/types";

const Content = () => {
  const [selectedTempl, setSelectedTempl] = useState<Template | null>(null);

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
        selectedTempl={selectedTempl}
        setSelectedTempl={setSelectedTempl}
      />
      <TemplatePreview selectedTempl={selectedTempl} />
    </Flex>
  );
};

export default Content;
