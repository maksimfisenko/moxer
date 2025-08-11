import { Button, Field, Fieldset, HStack, NumberInput } from "@chakra-ui/react";
import z from "zod";
import { Controller, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

const generateDataFormSchema = z.object({
  count: z.number().int().min(1).max(10),
});

type GenerateDataFormData = z.infer<typeof generateDataFormSchema>;

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
    <form onSubmit={handleSubmit(onFormSubmit)}>
      <Fieldset.Root w={"100%"} size={"lg"}>
        <HStack justify={"space-between"} w={"100%"}>
          <Fieldset.Legend>Generate from this template </Fieldset.Legend>
          <Fieldset.Content w={"50%"}>
            <Field.Root justifySelf={"center"}>
              <HStack justifyContent={"center"}>
                <Field.Label>Number of objects</Field.Label>

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
              </HStack>
            </Field.Root>
          </Fieldset.Content>
          <Button type="submit" alignSelf="flex-start" disabled={isLoading}>
            Submit
          </Button>
        </HStack>
      </Fieldset.Root>
    </form>
  );
};

export default GenerateDataForm;
