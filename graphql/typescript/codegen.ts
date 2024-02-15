import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  schema: "../api/*.graphql",
  documents: ["../api/documents/*.graphql"],
  generates: {
    "../../ui/src/components/gql/generated.ts": {
      plugins: [
        "typescript",
        "typescript-operations",
        "typescript-react-apollo",
      ],
      config: {
        withHooks: true,
        withComponent: false,
        withHOC: false,
        typeSuffix: "Dto",
      },
    },
  },
};
export default config;

