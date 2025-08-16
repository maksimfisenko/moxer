import { Toaster, toaster } from "@/components/ui/toaster";

interface ToasterPopup {
  type: string;
  title: string;
  description: string | undefined;
}

const ToasterPopup = ({ type, title, description }: ToasterPopup) => {
  return (
    <>
      {toaster.create({
        type: type,
        title: title,
        description: description,
      })}
      <Toaster />
    </>
  );
};

export default ToasterPopup;
