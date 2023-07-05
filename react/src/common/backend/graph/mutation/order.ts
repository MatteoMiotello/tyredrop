import {gql} from "../../../../__generated__";


export const NEW_ORDER = gql(`
    mutation createNewOrder( $userId: ID!, $userAddressId: ID! ) {
        newOrder(userId: $userId, userAddressId: $userAddressId) {
            id
        }
    }
`);
