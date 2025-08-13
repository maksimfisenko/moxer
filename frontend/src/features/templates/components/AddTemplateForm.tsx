import {
  Box,
  Button,
  Field,
  Fieldset,
  Flex,
  Input,
  Link,
  Stack,
} from "@chakra-ui/react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import z from "zod";
import type { CreateTemplateRequest } from "../types/types";
import MonacoEditor from "react-monaco-editor";
import * as monaco from "monaco-editor";

monaco.editor.defineTheme("customGray", {
  base: "vs",
  inherit: true,
  rules: [],
  colors: {
    "editor.background": "#f4f4f5",
  },
});

const addTemplateFormSchema = z.object({
  name: z.string().min(4, "Template name should be at least 4 characters long"),
  content: z
    .string()
    .refine(
      (val) => {
        try {
          JSON.parse(val);
          return true;
        } catch {
          return false;
        }
      },
      {
        message: "Invalid content format (should be JSON)",
      }
    )
    .transform((val) => JSON.parse(val)),
});

type AddTemplateFormData = z.infer<typeof addTemplateFormSchema>;

interface AddTemplateFormProps {
  isLoading: boolean;
  onFormSubmit: (req: CreateTemplateRequest) => void;
}

const AddTemplateForm = ({ isLoading, onFormSubmit }: AddTemplateFormProps) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
    watch,
  } = useForm<AddTemplateFormData>({
    resolver: zodResolver(addTemplateFormSchema),
  });

  return (
    <>
      <Flex
        flex={1}
        flexDirection={"column"}
        height={"100%"}
        align={"center"}
        justify={"center"}
        backgroundColor={"gray.100"}
      >
        <form onSubmit={handleSubmit(onFormSubmit)}>
          <Fieldset.Root
            size={"lg"}
            w={"lg"}
            backgroundColor={"white"}
            p={10}
            shadow={"md"}
            rounded={"lg"}
          >
            <Stack>
              <Fieldset.Legend fontSize={"xl"} alignSelf={"center"}>
                Create template
              </Fieldset.Legend>
              <Fieldset.HelperText alignSelf={"center"}>
                Please create a new template by filling the form below.
              </Fieldset.HelperText>
            </Stack>

            <Fieldset.Content>
              <Field.Root invalid={!!errors.name}>
                <Field.Label>
                  Name <Field.RequiredIndicator />
                </Field.Label>
                <Input
                  {...register("name")}
                  placeholder="my-template"
                  variant={"subtle"}
                />
                <Field.ErrorText>{errors.name?.message}</Field.ErrorText>
              </Field.Root>

              <Field.Root invalid={!!errors.content}>
                <Field.Label>
                  Content <Field.RequiredIndicator />
                </Field.Label>

                <Box
                  rounded="lg"
                  overflow="hidden"
                  w={"100%"}
                  p={5}
                  bgColor={"gray.100"}
                >
                  <MonacoEditor
                    height={"250px"}
                    language="json"
                    theme={"customGray"}
                    value={watch("content") || ""}
                    onChange={(value) =>
                      setValue("content", value || "", {
                        shouldValidate: false,
                      })
                    }
                    options={{
                      folding: false,
                      fontSize: 14,
                      minimap: { enabled: false },
                      formatOnPaste: true,
                      formatOnType: true,
                      lineNumbers: "off",
                      overviewRulerLanes: 0,
                      hideCursorInOverviewRuler: true,
                      scrollbar: {
                        verticalScrollbarSize: 0,
                        horizontalScrollbarSize: 0,
                      },
                      scrollBeyondLastLine: false,
                      lineDecorationsWidth: 0,
                      glyphMargin: false,
                      renderLineHighlight: "none",
                      guides: {
                        indentation: false,
                        highlightActiveIndentation: false,
                      },
                    }}
                  />
                </Box>
                <Field.HelperText>
                  <Link
                    target="_blank"
                    href="https://github.com/maksimfisenko/moxer?tab=readme-ov-file#sparkles-supported-template-variables"
                  >
                    Supported template variables
                  </Link>
                </Field.HelperText>

                <Field.ErrorText>"error"</Field.ErrorText>
              </Field.Root>
            </Fieldset.Content>

            <Button
              type="submit"
              disabled={isLoading}
              alignSelf={"center"}
              w={"3xs"}
              mb={2}
            >
              Submit
            </Button>
          </Fieldset.Root>
        </form>
      </Flex>
    </>
  );
};

export default AddTemplateForm;
