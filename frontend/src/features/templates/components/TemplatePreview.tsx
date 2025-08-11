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
  selectedTempl: Template | null;
}

const TemplatePreview = ({ selectedTempl }: TemplatePreviewProps) => {
  const { mutate, isPending } = useGenerateData();
  const [generatedData, setGeneratedData] = useState<GeneratedData | null>(
    null
  );

  const handleGenerateData = (data: GenerateDataRequest) => {
    mutate(
      {
        id: selectedTempl?.id || "",
        req: data,
      },
      {
        onSuccess: (data) => {
          console.log(data.data);
          setGeneratedData(data);
        },
        onError: (error: AxiosError<AxiosErrorResponseData>) => {
          console.log(error.response?.data.message);
        },
      }
    );
  };

  return (
    <VStack pl={4} align={"stretch"} flex={1}>
      {selectedTempl ? (
        <>
          <Heading size={"xl"} mb={2.5}>
            {selectedTempl.name}
          </Heading>

          <Separator />
          <TemplateInfo selectedTempl={selectedTempl} />
          <Separator />
          <GenerateDataForm
            isLoading={isPending}
            onFormSubmit={handleGenerateData}
          />
          <Separator />
          {generatedData ? (
            <TemplateTabs
              content={selectedTempl.content}
              generatedData={generatedData}
            />
          ) : (
            <TemplateTabs content={selectedTempl.content} />
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
