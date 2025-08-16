import MonacoEditor from "react-monaco-editor";
import * as monaco from "monaco-editor";
import jsonWorker from "monaco-editor/esm/vs/language/json/json.worker?worker";
import { Box } from "@chakra-ui/react";
import AdvisoryText from "@/components/ui/AdvisoryText";

self.MonacoEnvironment = {
  getWorker(_: any, __: any) {
    return new jsonWorker();
  },
};

monaco.languages.typescript.typescriptDefaults.setEagerModelSync(true);

monaco.editor.defineTheme("customGray", {
  base: "vs",
  inherit: true,
  rules: [],
  colors: {
    "editor.background": "#f4f4f5",
    "editor.lineHighlightBorder": "#f4f4f5",
    "editor.lineHighlightBackground": "#f4f4f5",
  },
});

interface JsonPreview {
  value: string | null | undefined;
  editorRef: any;
}

const JsonPreview = ({ value, editorRef }: JsonPreview) => {
  //   if (!value) {
  //     return (
  //       <Box bgColor={"gray.100"} height={"100%"} rounded={"lg"} p={4}>
  //         <AdvisoryText message="No generated data for this template. Generate data using a form on top." />
  //       </Box>
  //     );
  //   }

  return (
    <Box bgColor={"gray.100"} height={"100%"} rounded={"lg"} p={4}>
      {value ? (
        <MonacoEditor
          height={"100%"}
          language="json"
          theme={"customGray"}
          value={value}
          editorDidMount={(editor) => {
            editorRef.current = editor;
          }}
          options={{
            // Make read only without popup window
            readOnly: true,
            domReadOnly: true,
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
          }}
        />
      ) : (
        <AdvisoryText message="No generated data for this template. Generate data using a form on top." />
      )}
    </Box>
  );
};

export default JsonPreview;
