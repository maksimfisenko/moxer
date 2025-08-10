import { Box, Tabs } from "@chakra-ui/react";
import type { TemplateContent } from "../types/types";

interface TemplateContentPreviewProps {
  content: TemplateContent;
}

const TemplateContentPreview = ({ content }: TemplateContentPreviewProps) => {
  return (
    <Tabs.Root defaultValue="template" h={"100%"} variant={"enclosed"}>
      <Tabs.List>
        <Tabs.Trigger value="template">Template</Tabs.Trigger>
        <Tabs.Trigger value="data">Generated data</Tabs.Trigger>
      </Tabs.List>
      <Tabs.Content value="template" h={"88.5%"}>
        <Box as="pre" p={4} bg="gray.100" rounded="lg" fontSize="md" h={"100%"}>
          {JSON.stringify(content, null, 4)}
        </Box>
      </Tabs.Content>
    </Tabs.Root>
  );
};

export default TemplateContentPreview;
