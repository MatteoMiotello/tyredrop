import {CodegenConfig} from "@graphql-codegen/cli";

const config: CodegenConfig = {
    schema: 'http://localhost:8080/query',
    documents: [
        'src/**/*.tsx',
        "src/**/*.ts",
        "src/common/backend/graph/**/*.ts"
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