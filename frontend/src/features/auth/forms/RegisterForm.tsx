import {
  Button,
  Field,
  Fieldset,
  Flex,
  Input,
  Link,
  Stack,
} from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type { RegisterFormData, RegisterRequest } from "../types/authTypes";
import { registerFormSchema } from "../schemas/authSchemas";

interface RegisterFormProps {
  isLoading: boolean;
  onFormSubmit: (registerRequest: RegisterRequest) => void;
  onLinkClick: () => void;
}

const RegisterForm = ({
  isLoading,
  onFormSubmit,
  onLinkClick,
}: RegisterFormProps) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterFormData>({
    resolver: zodResolver(registerFormSchema),
  });

  return (
    <>
      <Flex
        flex={1}
        height={"100%"}
        align={"center"}
        justify={"center"}
        backgroundColor={"blue.50"}
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
              <Fieldset.Legend fontSize={"2xl"} alignSelf={"center"}>
                Sign up
              </Fieldset.Legend>
              <Fieldset.HelperText alignSelf={"center"}>
                Please create an account by filling the form below.
              </Fieldset.HelperText>
            </Stack>

            <Fieldset.Content>
              <Field.Root invalid={!!errors.email}>
                <Field.Label>
                  Email <Field.RequiredIndicator />
                </Field.Label>
                <Input
                  {...register("email")}
                  placeholder="email@example.com"
                  variant={"subtle"}
                />
                <Field.ErrorText>{errors.email?.message}</Field.ErrorText>
              </Field.Root>

              <Field.Root invalid={!!errors.password}>
                <Field.Label>
                  Password <Field.RequiredIndicator />
                </Field.Label>
                <PasswordInput
                  {...register("password")}
                  placeholder="********"
                  variant={"subtle"}
                />
                <Field.ErrorText>{errors.password?.message}</Field.ErrorText>
              </Field.Root>
            </Fieldset.Content>

            <Stack>
              <Button
                type="submit"
                alignSelf={"center"}
                w={"3xs"}
                mb={2}
                disabled={isLoading}
                variant={"subtle"}
              >
                Submit
              </Button>
              <Link alignSelf={"center"} onClick={onLinkClick}>
                Already have an account?
              </Link>
            </Stack>
          </Fieldset.Root>
        </form>
      </Flex>
    </>
  );
};

export default RegisterForm;
