import { Box } from "@chakra-ui/react";
import type { GeneratedData } from "../types/types";

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
      {JSON.stringify(data?.data, null, 4)}
    </Box>
  );
};

export default GeneratedDataPreview;
