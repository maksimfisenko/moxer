import { Flex, Heading, Separator, Text, VStack } from "@chakra-ui/react";
import type {
  GeneratedData,
  GenerateDataRequest,
  Template,
} from "../types/types";
import GenerateDataForm from "./GenerateDataForm";
import TemplateInfo from "./TemplateInfo";
import TemplateTabs from "./TemplateTabs";
import { useGenerateData } from "../hooks/use-generate-data";
import type { AxiosError } from "axios";
import type { AxiosErrorResponseData } from "@/types/types";
import { useState } from "react";

interface TemplatePreviewProps {
  template: Template | null;
}

const TemplatePreview = ({ template }: TemplatePreviewProps) => {
  const { mutate, isPending } = useGenerateData();
  const [data, setData] = useState<GeneratedData | null>(null);

  const handleGenerateData = (data: GenerateDataRequest) => {
    mutate(
      {
        id: template?.id || "",
        req: data,
      },
      {
        onSuccess: (data) => {
          setData(data);
        },
        onError: (error: AxiosError<AxiosErrorResponseData>) => {
          console.log(error.response?.data.message);
        },
      }
    );
  };

  return (
    <VStack pl={4} align={"stretch"} flex={1}>
      {template ? (
        <>
          <Heading size={"xl"} mb={2.5}>
            {template.name}
          </Heading>
          <Separator />
          <TemplateInfo selectedTempl={template} />
          <Separator />
          <GenerateDataForm
            isLoading={isPending}
            onFormSubmit={handleGenerateData}
          />
          <Separator />
          {data ? (
            <TemplateTabs content={template.content} generatedData={data} />
          ) : (
            <TemplateTabs content={template.content} />
          )}
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
