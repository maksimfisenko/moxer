import { Box, Spinner } from "@chakra-ui/react";
import { useGetTemplates } from "../hooks/use-get-templates";
import type { Template } from "../types/types";

interface TemplatesListProps {
  selectedTemplate: Template | null;
  setSelectedTemplate: React.Dispatch<React.SetStateAction<Template | null>>;
}

const TemplatesList = ({
  selectedTemplate,
  setSelectedTemplate,
}: TemplatesListProps) => {
  const { data, isLoading, isError, error } = useGetTemplates();

  if (isLoading) return <Spinner />;
  if (isError) return <div>{error.response?.data.message}</div>;
  if (!data) return null;

  return (
    <>
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
    </>
  );
};

export default TemplatesList;
