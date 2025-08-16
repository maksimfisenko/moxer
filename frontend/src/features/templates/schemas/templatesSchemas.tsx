import z from "zod";

export const generateDataFormSchema = z.object({
  count: z.number().int().min(1).max(10),
});

export const addTemplateFormSchema = z.object({
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
