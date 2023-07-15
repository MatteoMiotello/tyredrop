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
    "\n    fragment CartItems on CartResponse {\n        totalPrice {\n            value\n            currency {\n                iso_code\n                symbol\n                name\n            }\n        }\n        items {\n            id\n            quantity\n            productItem {\n                id\n                price {\n                    value\n                    currency {\n                        iso_code\n                    }\n                }\n                product {\n                    name\n                    code\n                    eprelProductCode\n                    brand {\n                        name\n                    }\n                }\n            }\n        }\n    }\n": types.CartItemsFragmentDoc,
    "\n    fragment UserAddressCollection on UserAddress {\n        ID\n        addressName\n        isDefault\n        addressLine1\n        addressLine2\n        city\n        country\n        postalCode\n        province\n    }\n": types.UserAddressCollectionFragmentDoc,
    "\n     \n    mutation addCart( $itemId: ID!, $quantity: Int) {\n        addItemToCart( itemId: $itemId, quantity: $quantity ) {\n           ...CartItems\n        }\n    } \n": types.AddCartDocument,
    "\n    \n    mutation editCart( $cartId: ID!, $quantity: Int! ) {\n        editCart( cartId: $cartId, quantity: $quantity ) {\n            ...CartItems\n        }\n    } \n": types.EditCartDocument,
    "\n    mutation CreateUserBilling( $input: CreateUserBilling! ) {\n        createUserBilling( billingInput: $input ) {\n            id\n            name\n            surname\n        }\n    }\n": types.CreateUserBillingDocument,
    "\n    mutation createNewOrder( $userId: ID!, $userAddressId: ID! ) {\n        newOrder(userId: $userId, userAddressId: $userAddressId) {\n            id\n        }\n    }\n": types.CreateNewOrderDocument,
    "\n    \n    mutation addAddress( $input: UserAddressInput! ) {\n        createUserAddress(userAddress: $input) {\n            ...UserAddressCollection\n        }\n    }\n": types.AddAddressDocument,
    "\n    \n    mutation editAddress( $id: ID!, $input: UserAddressInput! ) {\n        editUserAddress(id: $id, userAddress: $input) {\n            ...UserAddressCollection\n        }\n    }\n": types.EditAddressDocument,
    "\n    \n    mutation deleteAddress( $id: ID! ) {\n        deleteUserAddress(id: $id) {\n            ...UserAddressCollection\n        }\n    }\n": types.DeleteAddressDocument,
    "\n     query SearchBrands( $name: String! ) {\n        searchBrands(  name: $name ) {\n            id\n            name\n            code\n            quality\n        }\n     }  \n": types.SearchBrandsDocument,
    "\n    query userCarts {\n       carts {\n           totalPrice {\n               value\n               currency {\n                   iso_code\n                   symbol\n                   name\n               }\n           }\n           items {\n               id\n               quantity\n               productItem {\n                   id\n                   price {\n                       value\n                       currency {\n                           iso_code\n                       }\n                   }\n                   product {\n                       name\n                       code\n                       brand {\n                           name\n                       }\n                   }\n               }\n           }\n       }\n    }\n": types.UserCartsDocument,
    "\n    query GetLegalEntityTypes {\n        legalEntityTypes {\n            id\n            name\n            isPerson\n        }\n    }\n": types.GetLegalEntityTypesDocument,
    "\n    query fetchOrder( $orderId: ID! ) {\n        order(id: $orderId) {\n            id\n            currency {\n                iso_code\n                name\n            }\n            addressLine1\n            addressLine2\n            city\n            country\n            province\n            postalCode\n            status\n            createdAt\n            userBilling {\n                id\n                name\n                surname\n                vatNumber\n                fiscalCode\n                legalEntityType {\n                    name\n                }\n            }\n            orderRows {\n                id\n                amount\n                quantity\n                productItem {\n                    id\n                    price {\n                        value\n                        currency {\n                            iso_code\n                            symbol\n                        }\n                    }\n                    product {\n                        id\n                        name\n                    }\n                }\n            }\n        }\n    }\n": types.FetchOrderDocument,
    "\n    query fetchOrders( $userId: ID!, $pagination: PaginationInput ) {\n        userOrders(userId: $userId, pagination: $pagination) {\n            id\n            currency {\n                iso_code\n            }\n            status\n            addressLine1\n            addressLine2\n            city\n            province\n            postalCode\n            country\n            createdAt\n            orderRows {\n               amount\n            }\n        }\n    }\n": types.FetchOrdersDocument,
    "\n    query fetchUserAddresses( $userId: ID! ) {\n        userAddress( userId: $userId ) {\n            ID\n            addressName\n            isDefault\n            addressLine1\n            addressLine2\n            city\n            country\n            postalCode\n            province\n        }\n    }\n": types.FetchUserAddressesDocument,
    "\n    query fetchUserQuery( $userId: ID! ) {\n        userBilling( userId: $userId ) {\n            id\n            name\n            surname\n            addressLine1\n            addressLine2\n            city\n            country\n            province\n            cap\n            fiscalCode\n            vatNumber\n            sdiCode\n            sdiPec\n            legalEntityType {\n                name\n            }\n            user {\n                id\n                name\n                surname\n                email\n            }\n        }\n    }\n": types.FetchUserQueryDocument,
    "\n    query fetchUser( $userId: ID! ) {\n        user( ID: $userId ) {\n            id\n            name\n            surname\n            email\n            userBilling {\n                id\n                name\n                surname\n                addressLine1\n                addressLine2\n                city\n                country\n                province\n                cap\n                fiscalCode\n                vatNumber\n                sdiCode\n                sdiPec\n                legalEntityType {\n                    name\n                }\n            }\n        }\n    }\n": types.FetchUserDocument,
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
export function gql(source: "\n    fragment CartItems on CartResponse {\n        totalPrice {\n            value\n            currency {\n                iso_code\n                symbol\n                name\n            }\n        }\n        items {\n            id\n            quantity\n            productItem {\n                id\n                price {\n                    value\n                    currency {\n                        iso_code\n                    }\n                }\n                product {\n                    name\n                    code\n                    eprelProductCode\n                    brand {\n                        name\n                    }\n                }\n            }\n        }\n    }\n"): (typeof documents)["\n    fragment CartItems on CartResponse {\n        totalPrice {\n            value\n            currency {\n                iso_code\n                symbol\n                name\n            }\n        }\n        items {\n            id\n            quantity\n            productItem {\n                id\n                price {\n                    value\n                    currency {\n                        iso_code\n                    }\n                }\n                product {\n                    name\n                    code\n                    eprelProductCode\n                    brand {\n                        name\n                    }\n                }\n            }\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    fragment UserAddressCollection on UserAddress {\n        ID\n        addressName\n        isDefault\n        addressLine1\n        addressLine2\n        city\n        country\n        postalCode\n        province\n    }\n"): (typeof documents)["\n    fragment UserAddressCollection on UserAddress {\n        ID\n        addressName\n        isDefault\n        addressLine1\n        addressLine2\n        city\n        country\n        postalCode\n        province\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n     \n    mutation addCart( $itemId: ID!, $quantity: Int) {\n        addItemToCart( itemId: $itemId, quantity: $quantity ) {\n           ...CartItems\n        }\n    } \n"): (typeof documents)["\n     \n    mutation addCart( $itemId: ID!, $quantity: Int) {\n        addItemToCart( itemId: $itemId, quantity: $quantity ) {\n           ...CartItems\n        }\n    } \n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    \n    mutation editCart( $cartId: ID!, $quantity: Int! ) {\n        editCart( cartId: $cartId, quantity: $quantity ) {\n            ...CartItems\n        }\n    } \n"): (typeof documents)["\n    \n    mutation editCart( $cartId: ID!, $quantity: Int! ) {\n        editCart( cartId: $cartId, quantity: $quantity ) {\n            ...CartItems\n        }\n    } \n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    mutation CreateUserBilling( $input: CreateUserBilling! ) {\n        createUserBilling( billingInput: $input ) {\n            id\n            name\n            surname\n        }\n    }\n"): (typeof documents)["\n    mutation CreateUserBilling( $input: CreateUserBilling! ) {\n        createUserBilling( billingInput: $input ) {\n            id\n            name\n            surname\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    mutation createNewOrder( $userId: ID!, $userAddressId: ID! ) {\n        newOrder(userId: $userId, userAddressId: $userAddressId) {\n            id\n        }\n    }\n"): (typeof documents)["\n    mutation createNewOrder( $userId: ID!, $userAddressId: ID! ) {\n        newOrder(userId: $userId, userAddressId: $userAddressId) {\n            id\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    \n    mutation addAddress( $input: UserAddressInput! ) {\n        createUserAddress(userAddress: $input) {\n            ...UserAddressCollection\n        }\n    }\n"): (typeof documents)["\n    \n    mutation addAddress( $input: UserAddressInput! ) {\n        createUserAddress(userAddress: $input) {\n            ...UserAddressCollection\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    \n    mutation editAddress( $id: ID!, $input: UserAddressInput! ) {\n        editUserAddress(id: $id, userAddress: $input) {\n            ...UserAddressCollection\n        }\n    }\n"): (typeof documents)["\n    \n    mutation editAddress( $id: ID!, $input: UserAddressInput! ) {\n        editUserAddress(id: $id, userAddress: $input) {\n            ...UserAddressCollection\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    \n    mutation deleteAddress( $id: ID! ) {\n        deleteUserAddress(id: $id) {\n            ...UserAddressCollection\n        }\n    }\n"): (typeof documents)["\n    \n    mutation deleteAddress( $id: ID! ) {\n        deleteUserAddress(id: $id) {\n            ...UserAddressCollection\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n     query SearchBrands( $name: String! ) {\n        searchBrands(  name: $name ) {\n            id\n            name\n            code\n            quality\n        }\n     }  \n"): (typeof documents)["\n     query SearchBrands( $name: String! ) {\n        searchBrands(  name: $name ) {\n            id\n            name\n            code\n            quality\n        }\n     }  \n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query userCarts {\n       carts {\n           totalPrice {\n               value\n               currency {\n                   iso_code\n                   symbol\n                   name\n               }\n           }\n           items {\n               id\n               quantity\n               productItem {\n                   id\n                   price {\n                       value\n                       currency {\n                           iso_code\n                       }\n                   }\n                   product {\n                       name\n                       code\n                       brand {\n                           name\n                       }\n                   }\n               }\n           }\n       }\n    }\n"): (typeof documents)["\n    query userCarts {\n       carts {\n           totalPrice {\n               value\n               currency {\n                   iso_code\n                   symbol\n                   name\n               }\n           }\n           items {\n               id\n               quantity\n               productItem {\n                   id\n                   price {\n                       value\n                       currency {\n                           iso_code\n                       }\n                   }\n                   product {\n                       name\n                       code\n                       brand {\n                           name\n                       }\n                   }\n               }\n           }\n       }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query GetLegalEntityTypes {\n        legalEntityTypes {\n            id\n            name\n            isPerson\n        }\n    }\n"): (typeof documents)["\n    query GetLegalEntityTypes {\n        legalEntityTypes {\n            id\n            name\n            isPerson\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query fetchOrder( $orderId: ID! ) {\n        order(id: $orderId) {\n            id\n            currency {\n                iso_code\n                name\n            }\n            addressLine1\n            addressLine2\n            city\n            country\n            province\n            postalCode\n            status\n            createdAt\n            userBilling {\n                id\n                name\n                surname\n                vatNumber\n                fiscalCode\n                legalEntityType {\n                    name\n                }\n            }\n            orderRows {\n                id\n                amount\n                quantity\n                productItem {\n                    id\n                    price {\n                        value\n                        currency {\n                            iso_code\n                            symbol\n                        }\n                    }\n                    product {\n                        id\n                        name\n                    }\n                }\n            }\n        }\n    }\n"): (typeof documents)["\n    query fetchOrder( $orderId: ID! ) {\n        order(id: $orderId) {\n            id\n            currency {\n                iso_code\n                name\n            }\n            addressLine1\n            addressLine2\n            city\n            country\n            province\n            postalCode\n            status\n            createdAt\n            userBilling {\n                id\n                name\n                surname\n                vatNumber\n                fiscalCode\n                legalEntityType {\n                    name\n                }\n            }\n            orderRows {\n                id\n                amount\n                quantity\n                productItem {\n                    id\n                    price {\n                        value\n                        currency {\n                            iso_code\n                            symbol\n                        }\n                    }\n                    product {\n                        id\n                        name\n                    }\n                }\n            }\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query fetchOrders( $userId: ID!, $pagination: PaginationInput ) {\n        userOrders(userId: $userId, pagination: $pagination) {\n            id\n            currency {\n                iso_code\n            }\n            status\n            addressLine1\n            addressLine2\n            city\n            province\n            postalCode\n            country\n            createdAt\n            orderRows {\n               amount\n            }\n        }\n    }\n"): (typeof documents)["\n    query fetchOrders( $userId: ID!, $pagination: PaginationInput ) {\n        userOrders(userId: $userId, pagination: $pagination) {\n            id\n            currency {\n                iso_code\n            }\n            status\n            addressLine1\n            addressLine2\n            city\n            province\n            postalCode\n            country\n            createdAt\n            orderRows {\n               amount\n            }\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query fetchUserAddresses( $userId: ID! ) {\n        userAddress( userId: $userId ) {\n            ID\n            addressName\n            isDefault\n            addressLine1\n            addressLine2\n            city\n            country\n            postalCode\n            province\n        }\n    }\n"): (typeof documents)["\n    query fetchUserAddresses( $userId: ID! ) {\n        userAddress( userId: $userId ) {\n            ID\n            addressName\n            isDefault\n            addressLine1\n            addressLine2\n            city\n            country\n            postalCode\n            province\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query fetchUserQuery( $userId: ID! ) {\n        userBilling( userId: $userId ) {\n            id\n            name\n            surname\n            addressLine1\n            addressLine2\n            city\n            country\n            province\n            cap\n            fiscalCode\n            vatNumber\n            sdiCode\n            sdiPec\n            legalEntityType {\n                name\n            }\n            user {\n                id\n                name\n                surname\n                email\n            }\n        }\n    }\n"): (typeof documents)["\n    query fetchUserQuery( $userId: ID! ) {\n        userBilling( userId: $userId ) {\n            id\n            name\n            surname\n            addressLine1\n            addressLine2\n            city\n            country\n            province\n            cap\n            fiscalCode\n            vatNumber\n            sdiCode\n            sdiPec\n            legalEntityType {\n                name\n            }\n            user {\n                id\n                name\n                surname\n                email\n            }\n        }\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    query fetchUser( $userId: ID! ) {\n        user( ID: $userId ) {\n            id\n            name\n            surname\n            email\n            userBilling {\n                id\n                name\n                surname\n                addressLine1\n                addressLine2\n                city\n                country\n                province\n                cap\n                fiscalCode\n                vatNumber\n                sdiCode\n                sdiPec\n                legalEntityType {\n                    name\n                }\n            }\n        }\n    }\n"): (typeof documents)["\n    query fetchUser( $userId: ID! ) {\n        user( ID: $userId ) {\n            id\n            name\n            surname\n            email\n            userBilling {\n                id\n                name\n                surname\n                addressLine1\n                addressLine2\n                city\n                country\n                province\n                cap\n                fiscalCode\n                vatNumber\n                sdiCode\n                sdiPec\n                legalEntityType {\n                    name\n                }\n            }\n        }\n    }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;