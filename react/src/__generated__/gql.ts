/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n    mutation CreateUserBilling( $input: CreateUserBilling! ) {\n        createUserBilling( billingInput: $input ) {\n            id\n            name\n            surname\n        }\n    }\n": types.CreateUserBillingDocument,
    "\n     query SearchBrands( $name: String! ) {\n        searchBrands(  name: $name ) {\n            id\n            name\n        }\n     }  \n": types.SearchBrandsDocument,
    "\n    query GetLegalEntityTypes {\n        legalEntityTypes {\n            id\n            name\n            isPerson\n        }\n    }\n": types.GetLegalEntityTypesDocument,
    "\n    query getAllCategories {\n        productCategories {\n            id\n            name\n            code\n            specifications {\n                id\n                code\n                name\n                type\n                mandatory\n                searchable\n            }\n        }\n    }\n": types.GetAllCategoriesDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    mutation CreateUserBilling( $input: CreateUserBilling! ) {\n        createUserBilling( billingInput: $input ) {\n            id\n            name\n            surname\n        }\n    }\n"): (typeof documents)["\n    mutation CreateUserBilling( $input: CreateUserBilling! ) {\n        createUserBilling( billingInput: $input ) {\n            id\n            name\n            surname\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n     query SearchBrands( $name: String! ) {\n        searchBrands(  name: $name ) {\n            id\n            name\n        }\n     }  \n"): (typeof documents)["\n     query SearchBrands( $name: String! ) {\n        searchBrands(  name: $name ) {\n            id\n            name\n        }\n     }  \n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query GetLegalEntityTypes {\n        legalEntityTypes {\n            id\n            name\n            isPerson\n        }\n    }\n"): (typeof documents)["\n    query GetLegalEntityTypes {\n        legalEntityTypes {\n            id\n            name\n            isPerson\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query getAllCategories {\n        productCategories {\n            id\n            name\n            code\n            specifications {\n                id\n                code\n                name\n                type\n                mandatory\n                searchable\n            }\n        }\n    }\n"): (typeof documents)["\n    query getAllCategories {\n        productCategories {\n            id\n            name\n            code\n            specifications {\n                id\n                code\n                name\n                type\n                mandatory\n                searchable\n            }\n        }\n    }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;