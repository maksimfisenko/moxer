import { Box } from "@chakra-ui/react";
import type { TemplateContent } from "../types/types";
import Json from "./Json";

interface ContentPreviewProps {
  content: TemplateContent;
}

const ContentPreview = ({ content }: ContentPreviewProps) => {
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
      maxH={"345px"}
      overflowY={"auto"}
    >
      <Json
        height={"100%"}
        readOnly={true}
        value={JSON.stringify(content, null, 4)}
      />
    </Box>
  );
};

export default ContentPreview;
