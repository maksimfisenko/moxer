import { Tabs } from "@chakra-ui/react";
import type { GeneratedData, Template } from "../types/types";
import GeneratedDataPreview from "./GeneratedDataPreview";
import ContentPreview from "./ContentPreview";
import { useEffect, useState } from "react";

interface TemplateTabsProps {
  selectedTemplate: Template;
  generatedData?: GeneratedData;
}

const TemplateTabs = ({
  selectedTemplate,
  generatedData,
}: TemplateTabsProps) => {
  const [activeTab, setActiveTab] = useState<string>("template");

  useEffect(() => {
    setActiveTab("template");
  }, [selectedTemplate]);

  return (
    <Tabs.Root
      flex={1}
      flexDirection={"column"}
      display={"flex"}
      variant={"enclosed"}
      value={activeTab}
      onValueChange={(e) => setActiveTab(e.value)}
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
        <ContentPreview content={selectedTemplate.content} />
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
