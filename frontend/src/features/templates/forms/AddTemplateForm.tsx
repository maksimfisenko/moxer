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
import type {
  AddTemplateFormData,
  CreateTemplateRequest,
} from "../types/templatesTypes";
import { addTemplateFormSchema } from "../schemas/templatesSchemas";
import JsonEditor from "@/shared/components/JsonEditor";

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
                  placeholder="Template name"
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
                  bgColor={"blue.50"}
                >
                  <JsonEditor
                    readOnly={false}
                    height={"250px"}
                    value={watch("content") || ""}
                    onChange={(value) =>
                      setValue("content", value || "", {
                        shouldValidate: false,
                      })
                    }
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
              variant={"subtle"}
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
