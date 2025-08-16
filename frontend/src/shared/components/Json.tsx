import MonacoEditor, { type ChangeHandler } from "react-monaco-editor";
import * as monaco from "monaco-editor";
import jsonWorker from "monaco-editor/esm/vs/language/json/json.worker?worker";

self.MonacoEnvironment = {
  getWorker(_: any, __: any) {
    return new jsonWorker();
  },
};

monaco.languages.typescript.typescriptDefaults.setEagerModelSync(true);

monaco.editor.defineTheme("customGray1", {
  base: "vs",
  inherit: true,
  rules: [],
  colors: {
    "editor.background": "#f4f4f5",
    "editor.lineHighlightBorder": "#f4f4f5",
    "editor.lineHighlightBackground": "#f4f4f5",
  },
});

interface JsonProps {
  readOnly: boolean;
  value?: string | null | undefined;
  onChange?: ChangeHandler | undefined;
  height: string | number | undefined;
}

const Json = ({ readOnly, value, onChange, height }: JsonProps) => {
  return (
    <MonacoEditor
      height={height}
      language="json"
      theme={"customGray1"}
      value={value}
      onChange={onChange}
      options={{
        folding: false,
        fontSize: 14,
        minimap: { enabled: false },
        formatOnPaste: true,
        formatOnType: true,
        lineNumbers: "off",
        overviewRulerLanes: 0,
        hideCursorInOverviewRuler: true,
        scrollbar: {
          verticalScrollbarSize: 0,
          horizontalScrollbarSize: 0,
        },
        scrollBeyondLastLine: false,
        lineDecorationsWidth: 0,
        glyphMargin: false,
        renderLineHighlight: "none",
        guides: {
          indentation: false,
          highlightActiveIndentation: false,
        },
        readOnly: readOnly,
        domReadOnly: true,
        renderValidationDecorations: "off",
        matchBrackets: "never",
        occurrencesHighlight: "off",
        selectionHighlight: false,
      }}
    />
  );
};

export default Json;
