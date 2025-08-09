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

const RegisterForm = () => {
  return (
    <Flex
      flex={1}
      align={"center"}
      justify={"center"}
      backgroundColor={"gray.100"}
    >
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
          <Field.Root required>
            <Field.Label>
              Email <Field.RequiredIndicator />
            </Field.Label>
            <Input
              name="email"
              placeholder="email@example.com"
              variant={"subtle"}
            />
          </Field.Root>

          <Field.Root required>
            <Field.Label>
              Password <Field.RequiredIndicator />
            </Field.Label>
            <PasswordInput
              name="password"
              placeholder="********"
              variant={"subtle"}
            />
          </Field.Root>
        </Fieldset.Content>

        <Stack>
          <Button type="submit" alignSelf={"center"} w={"3xs"} mb={2}>
            Submit
          </Button>
          <Link alignSelf={"center"}>Already have an account?</Link>
        </Stack>
      </Fieldset.Root>
    </Flex>
  );
};

export default RegisterForm;
