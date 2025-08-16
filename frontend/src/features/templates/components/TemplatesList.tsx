import LoadingSpinner from "@/shared/components/LoadingSpinner";
import { Box } from "@chakra-ui/react";
import ErrorText from "@/shared/components/ErrorText";
import AdvisoryText from "@/shared/components/AdvisoryText";
import type { Template } from "../types/templatesTypes";
import { useGetTemplates } from "../hooks/useGetTemplate";

interface TemplatesListProps {
  selectedTemplate: Template | null;
  setSelectedTemplate: React.Dispatch<React.SetStateAction<Template | null>>;
}

const TemplatesList = ({
  selectedTemplate,
  setSelectedTemplate,
}: TemplatesListProps) => {
  const { data, isLoading, isError, error } = useGetTemplates();

  if (isLoading) return <LoadingSpinner />;

  if (isError)
    return (
      <ErrorText
        message="Failed to load your templates"
        details={error.response?.data.message}
      />
    );

  if (!data) return null;

  if (data.length == 0)
    return (
      <AdvisoryText message="You don't have any saved templates yet. Create a new one by clicking the green button on top." />
    );

  return (
    <Box h={"480px"} overflowY={"auto"}>
      {data.map((templ) => (
        <Box
          key={templ.id}
          p={2}
          cursor={"pointer"}
          _hover={{ bg: "gray.50" }}
          bg={selectedTemplate?.id === templ.id ? "gray.100" : "transparent"}
          rounded={"md"}
          onClick={() => setSelectedTemplate(templ)}
        >
          {templ.name}
        </Box>
      ))}
    </Box>
  );
};

export default TemplatesList;
