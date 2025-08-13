import { Box } from "@chakra-ui/react";
import type { TemplateContent } from "../types/types";
import JsonView from "@uiw/react-json-view";

interface JsonPreviewProps {
  content: TemplateContent;
}

const JsonPreview = ({ content }: JsonPreviewProps) => {
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
      <JsonView
        value={content}
        indentWidth={16}
        displayDataTypes={false}
        collapsed={false}
        enableClipboard={false}
        displayObjectSize={false}
        shortenTextAfterLength={100}
        style={{ fontSize: "14px" }}
      />
    </Box>
  );
};

export default JsonPreview;
