import { Tabs } from "@chakra-ui/react";
import type { GeneratedData, TemplateContent } from "../types/types";
import TemplateContentPreview from "./TemplateContentPreview";
import GeneratedDataPreview from "./GeneratedDataPreview";

interface TemplateTabsProps {
  content: TemplateContent;
  generatedData?: GeneratedData;
}

const TemplateTabs = ({ content, generatedData }: TemplateTabsProps) => {
  return (
    <Tabs.Root
      flex={1}
      flexDirection={"column"}
      display={"flex"}
      defaultValue="template"
      variant={"enclosed"}
    >
      <Tabs.List>
        <Tabs.Trigger value="template">Template</Tabs.Trigger>
        <Tabs.Trigger value="data">Generated data</Tabs.Trigger>
      </Tabs.List>

      <Tabs.Content
        value="template"
        display={"flex"}
        flex={1}
        flexDirection={"column"}
      >
        <TemplateContentPreview content={content} />
      </Tabs.Content>
      <Tabs.Content
        value="data"
        display={"flex"}
        flex={1}
        flexDirection={"column"}
      >
        <GeneratedDataPreview data={generatedData} />
      </Tabs.Content>
    </Tabs.Root>
  );
};

export default TemplateTabs;
