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
import z from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Toaster } from "@/components/ui/toaster";
import type { RegisterRequest } from "../types/types";

const registerFormSchema = z.object({
  email: z.email("Invalid email address"),
  password: z.string().min(8, "Password should be at least 8 characters long"),
});

type RegisterFormData = z.infer<typeof registerFormSchema>;

interface RegisterFormProps {
  isLoading: boolean;
  onFormSubmit: (registerRequest: RegisterRequest) => void;
}

const RegisterForm = ({ isLoading, onFormSubmit }: RegisterFormProps) => {
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
              >
                Submit
              </Button>
              <Link alignSelf={"center"}>Already have an account?</Link>
            </Stack>
          </Fieldset.Root>
        </form>
      </Flex>
      <Toaster />
    </>
  );
};

export default RegisterForm;
