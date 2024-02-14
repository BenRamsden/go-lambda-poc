import type { CodegenConfig } from '@graphql-codegen/cli'
 
const config: CodegenConfig = {
   schema: '../api/*.graphql',
   documents: [],
   generates: {
      '../../ui/src/components/gql/generated/': {
        preset: 'client',
      }
   }
}
export default config