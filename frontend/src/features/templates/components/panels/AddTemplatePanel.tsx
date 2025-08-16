import AddTemplateButton from "../AddTemplateButton";
import AddTemplateDrawer from "../AddTemplateDrawer";

interface AddTemplatePanelProps {
  drawerIsOpen: boolean;
  setDrawerOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const AddTemplatePanel = ({
  drawerIsOpen,
  setDrawerOpen,
}: AddTemplatePanelProps) => {
  return (
    <>
      <AddTemplateButton setDrawerOpen={setDrawerOpen} />
      <AddTemplateDrawer isOpen={drawerIsOpen} setOpen={setDrawerOpen} />
    </>
  );
};

export default AddTemplatePanel;
