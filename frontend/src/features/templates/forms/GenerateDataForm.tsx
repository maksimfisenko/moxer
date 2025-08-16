import {
  Button,
  Field,
  Fieldset,
  Flex,
  NumberInput,
  Stack,
} from "@chakra-ui/react";
import { Controller, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type { GenerateDataFormData } from "../types/templatesTypes";
import { generateDataFormSchema } from "../schemas/templatesSchemas";

interface GenerateDataFormProps {
  isLoading: boolean;
  onFormSubmit: (data: GenerateDataFormData) => void;
}

const GenerateDataForm = ({
  isLoading,
  onFormSubmit,
}: GenerateDataFormProps) => {
  const { control, handleSubmit } = useForm<GenerateDataFormData>({
    resolver: zodResolver(generateDataFormSchema),
    defaultValues: { count: 3 },
  });

  return (
    <Flex
      height={"100%"}
      align={"center"}
      justify={"center"}
      border={"1px solid gray"}
      rounded={"lg"}
      width={"30%"}
    >
      <Fieldset.Root size={"lg"} w="sm" p={4} rounded={"lg"}>
        <form onSubmit={handleSubmit(onFormSubmit)}>
          <Stack>
            <Fieldset.Legend fontSize={"lg"} alignSelf={"center"}>
              Generate from this template
            </Fieldset.Legend>
            <Fieldset.Content>
              <Field.Root alignItems={"center"}>
                <Field.Label>Number of objects (1-10)</Field.Label>

                <Controller
                  name="count"
                  control={control}
                  render={({ field }) => (
                    <NumberInput.Root
                      w={"50%"}
                      defaultValue="3"
                      min={1}
                      max={10}
                      value={field.value.toString()}
                      onValueChange={(val) => {
                        field.onChange(Number(val.value));
                      }}
                    >
                      <NumberInput.Control />
                      <NumberInput.Input />
                    </NumberInput.Root>
                  )}
                />
              </Field.Root>
            </Fieldset.Content>

            <Button
              width={"50%"}
              alignSelf={"center"}
              type="submit"
              variant={"subtle"}
              disabled={isLoading}
            >
              Submit
            </Button>
          </Stack>
        </form>
      </Fieldset.Root>
    </Flex>
  );
};

export default GenerateDataForm;
