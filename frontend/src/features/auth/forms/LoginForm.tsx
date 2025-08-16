import { Button, Field, Fieldset, Input, Link, Stack } from "@chakra-ui/react";
import { PasswordInput } from "@/components/ui/password-input";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { loginFormSchema } from "../schemas/authSchemas";
import type { LoginFormData, LoginRequest } from "../types/authTypes";

interface LoginFormProps {
  isLoading: boolean;
  onFormSubmit: (loginRequest: LoginRequest) => void;
  onLinkClick: () => void;
}

const LoginForm = ({
  isLoading,
  onFormSubmit,
  onLinkClick,
}: LoginFormProps) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormData>({
    resolver: zodResolver(loginFormSchema),
  });

  return (
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
            Sign in
          </Fieldset.Legend>
          <Fieldset.HelperText alignSelf={"center"}>
            Please sign in your account by filling the form below.
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
            Don't have an acccount?
          </Link>
        </Stack>
      </Fieldset.Root>
    </form>
  );
};

export default LoginForm;
