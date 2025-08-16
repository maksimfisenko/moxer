import MonacoEditor, { type ChangeHandler } from "react-monaco-editor";
import * as monaco from "monaco-editor";
import jsonWorker from "monaco-editor/esm/vs/language/json/json.worker?worker";

self.MonacoEnvironment = {
  getWorker(_: any, __: any) {
    return new jsonWorker();
  },
};

monaco.languages.typescript.typescriptDefaults.setEagerModelSync(true);

monaco.editor.defineTheme("customBlue", {
  base: "vs",
  inherit: true,
  rules: [],
  colors: {
    "editor.background": "#eff6ff",
    "editor.lineHighlightBorder": "#eff6ff",
    "editor.lineHighlightBackground": "#eff6ff",
  },
});

interface JsonEditorProps {
  value: string | null;
  height: string;
  readOnly: boolean;
  onChange?: ChangeHandler | undefined;
}

const JsonEditor = ({ value, height, readOnly, onChange }: JsonEditorProps) => {
  return (
    <MonacoEditor
      height={height}
      language="json"
      theme={"customBlue"}
      value={value}
      onChange={onChange}
      options={{
        // Make read only without popup window
        readOnly: readOnly,
        domReadOnly: readOnly,
        // Hide minimap
        minimap: { enabled: false },
        // Hide line numbers
        lineNumbers: "off",
        lineDecorationsWidth: 0,
        // Hide mathing bracket selection
        matchBrackets: "never",
        // Hide line hightlight
        renderLineHighlight: "none",
        // Hide occurencies highlight
        occurrencesHighlight: "off",
        selectionHighlight: false,
        // Hide folding
        folding: false,
        // Hide not needed scroll
        scrollBeyondLastLine: false,
        // Set font size
        fontSize: 14,
        // Hide guide lines
        guides: {
          indentation: false,
        },
        // Hide scrollbar not needed elements
        overviewRulerLanes: 0,
        scrollbar: {
          useShadows: false,
        },
        stickyScroll: {
          enabled: false,
        },
        // Format while typing
        formatOnPaste: true,
        formatOnType: true,
      }}
    />
  );
};

export default JsonEditor;
