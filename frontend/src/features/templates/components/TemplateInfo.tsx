import { DataList } from "@chakra-ui/react";
import type { Template } from "../types/types";

interface TemplateInfoProps {
  selectedTemplate: Template;
}

const TemplateInfo = ({ selectedTemplate }: TemplateInfoProps) => {
  return (
    <DataList.Root orientation={"horizontal"}>
      <DataList.Item>
        <DataList.ItemLabel>Created at</DataList.ItemLabel>
        <DataList.ItemValue>
          {selectedTemplate.created_at.toLocaleString("ru-RU")}
        </DataList.ItemValue>
      </DataList.Item>

      <DataList.Item>
        <DataList.ItemLabel>Updated at</DataList.ItemLabel>
        <DataList.ItemValue>
          {selectedTemplate.updated_at.toLocaleString("ru-RU")}
        </DataList.ItemValue>
      </DataList.Item>
    </DataList.Root>
  );
};

export default TemplateInfo;
