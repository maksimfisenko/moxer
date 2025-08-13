import { Box } from "@chakra-ui/react";
import type { GeneratedData } from "../types/types";
import JsonView from "@uiw/react-json-view";

interface GeneratedDataPreviewProps {
  data?: GeneratedData;
}

const GeneratedDataPreview = ({ data }: GeneratedDataPreviewProps) => {
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
      {data && (
        <JsonView
          value={data?.data}
          indentWidth={16}
          displayDataTypes={false}
          collapsed={false}
          enableClipboard={false}
          displayObjectSize={false}
          shortenTextAfterLength={100}
          style={{ fontSize: "14px" }}
        />
      )}
    </Box>
  );
};

export default GeneratedDataPreview;
