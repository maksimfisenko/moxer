import { Tabs } from "@chakra-ui/react";
import type { GeneratedData, Template } from "../types/types";
import { useEffect, useRef, useState } from "react";
import JsonPreview from "./JsonPreview";
import type { monaco } from "react-monaco-editor";

interface TemplateTabsProps {
  selectedTemplate: Template;
  generatedData?: GeneratedData;
}

const TemplateTabs = ({
  selectedTemplate,
  generatedData,
}: TemplateTabsProps) => {
  const [activeTab, setActiveTab] = useState<string>("template");
  const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);

  useEffect(() => {
    setActiveTab("template");
  }, [selectedTemplate]);

  useEffect(() => {
    if (activeTab === "data" && editorRef.current) {
      setTimeout(() => editorRef.current?.layout(), 0);
    }
  }, [activeTab, generatedData]);

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
        hidden={activeTab !== "template"}
        value="template"
        display={"flex"}
        flex={1}
        flexDirection={"column"}
      >
        <JsonPreview
          value={JSON.stringify(selectedTemplate.content, null, 4)}
          editorRef={editorRef}
        />
      </Tabs.Content>
      <Tabs.Content
        hidden={activeTab !== "data"}
        value="data"
        display={"flex"}
        flex={1}
        flexDirection={"column"}
      >
        <JsonPreview
          value={JSON.stringify(generatedData?.data, null, 4)}
          editorRef={editorRef}
        />
      </Tabs.Content>
    </Tabs.Root>
  );
};

export default TemplateTabs;
