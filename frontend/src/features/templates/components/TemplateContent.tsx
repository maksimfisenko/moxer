import JsonEditor from "@/shared/components/JsonEditor";
import { Box } from "@chakra-ui/react";

interface TemplateContentProps {
  value: string | null;
}

const TemplateContent = ({ value }: TemplateContentProps) => {
  return (
    <Box bgColor={"blue.50"} height={"100%"} rounded={"lg"} p={2} flex={8}>
      <JsonEditor height="100%" readOnly={true} value={value} />
    </Box>
  );
};

export default TemplateContent;
