import { HStack, VStack } from "@chakra-ui/react";
import type {
  GeneratedData,
  GenerateDataRequest,
  Template,
} from "../../types/templatesTypes";
import GenerateDataForm from "../../forms/GenerateDataForm";
import { useGenerateData } from "../../hooks/useGenerateData";
import type { AxiosError } from "axios";
import { useEffect, useState } from "react";
import AdvisoryText from "@/shared/components/AdvisoryText";
import SeparatorWithText from "../../../../shared/components/SeparatorWithText";
import TemplateInfo from "../TemplateInfo";
import TemplateContent from "../TemplateContent";
import TemplateGeneratedData from "../TemplateGeneratedData";

interface TemplatePanelProps {
  selectedTemplate: Template | null;
}

const TemplatePanel = ({ selectedTemplate }: TemplatePanelProps) => {
  const { mutate, isPending } = useGenerateData();
  const [generatedData, setGeneratedData] = useState<GeneratedData | null>(
    null
  );

  useEffect(() => setGeneratedData(null), [selectedTemplate]);

  const handleGenerateData = (data: GenerateDataRequest) => {
    mutate(
      {
        id: selectedTemplate?.id || "",
        req: data,
      },
      {
        onSuccess: (data) => {
          setGeneratedData(data);
        },
        onError: (error: AxiosError<any>) => {
          console.log(error.response?.data.message);
        },
      }
    );
  };

  if (!selectedTemplate) {
    return (
      <VStack pl={4} align={"stretch"} flex={1} height={"100%"}>
        <AdvisoryText message="Choose the template from the left panel or create a new one." />
      </VStack>
    );
  }

  return (
    <VStack pl={4} align={"stretch"} flex={1} height={"100%"}>
      <TemplateInfo selectedTemplate={selectedTemplate} />
      <SeparatorWithText text="Template information" />
      <HStack height={"70%"} align={"stretch"}>
        <TemplateContent
          value={JSON.stringify(selectedTemplate.content, null, 4)}
        />
        <GenerateDataForm
          isLoading={isPending}
          onFormSubmit={handleGenerateData}
        />
      </HStack>
      <SeparatorWithText text="Generated data" />
      <TemplateGeneratedData
        value={JSON.stringify(generatedData?.data, null, 4)}
      />
    </VStack>
  );
};

export default TemplatePanel;
