import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  schema: "../api/*.graphql",
  documents: [],
  config: {
    withHooks: true,
    withComponent: false,
    withHOC: false,
    typeSuffix: "Dto",
  },
  generates: {
    "../../ui/src/components/gql/generated/types.ts": {
      plugins: [
        "typescript",
        "typescript-operations",
        "typescript-react-apollo",
      ],
    },
  },
};
export default config;

