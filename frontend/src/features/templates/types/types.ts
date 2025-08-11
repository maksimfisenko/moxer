export type TemplateContent = Record<string, unknown>;

export interface GeneratedData {
  data: Record<string, unknown>[];
}

export interface Template {
  id: string;
  name: string;
  content: TemplateContent;
  created_at: Date;
  updated_at: Date;
}

export interface GenerateDataRequest {
  count: number;
}

export interface CreateTemplateRequest {
  name: string;
  content: any;
}
