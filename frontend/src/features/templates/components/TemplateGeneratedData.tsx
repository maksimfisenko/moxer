import { Box } from "@chakra-ui/react";
import AdvisoryText from "@/shared/components/AdvisoryText";
import JsonEditor from "@/shared/components/JsonEditor";

interface TemplateGeneratedDataProps {
  value: string | null;
}

const TemplateGeneratedData = ({ value }: TemplateGeneratedDataProps) => {
  return (
    <Box bgColor={"blue.50"} height={"100%"} rounded={"lg"} p={2}>
      {value ? (
        <JsonEditor height="100%" readOnly={true} value={value} />
      ) : (
        <AdvisoryText message="Generate data for this template using the form on top." />
      )}
    </Box>
  );
};

export default TemplateGeneratedData;
