import {CodegenConfig} from "@graphql-codegen/cli";
import backend from "./src/config/backend";

const config: CodegenConfig = {
    schema: backend.graphEndpoint,
    documents: [
        'src/**/*.tsx',
        "src/**/*.ts"
    ],
    generates: {
        './src/__generated__/': {
            preset: 'client',
            plugins: [],
            presetConfig: {
                gqlTagName: 'gql',
            }
        }
    },
    ignoreNoDocuments: true,
};

export default config;