import { Box } from "@chakra-ui/react";
import type { TemplateContent } from "../types/types";

interface TemplateContentPreviewProps {
  content: TemplateContent;
}

const TemplateContentPreview = ({ content }: TemplateContentPreviewProps) => {
  return (
    <Box
      display={"flex"}
      flex={1}
      as="pre"
      p={4}
      bg="gray.100"
      rounded="lg"
      fontSize="md"
      whiteSpace="pre-wrap"
      wordBreak="break-word"
      overflowWrap="break-word"
    >
      {JSON.stringify(content, null, 4)}
    </Box>
  );
};

export default TemplateContentPreview;
