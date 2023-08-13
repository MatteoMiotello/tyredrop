import {gql} from "@apollo/client";


export const NEW_ORDER = gql(`
    mutation createNewOrder( $userId: ID!, $userAddressId: ID! ) {
        newOrder(userId: $userId, userAddressId: $userAddressId) {
            id
        }
    }
`);

export const CONFIRM_ORDER = gql(`
    mutation confirmOrder( $orderId: ID!) {
        confirmOrder(orderID: $orderId) {
            id
        }
    }
` );

export const PAY_ORDER = gql(`
    mutation payOrder( $orderId: ID!, $methodCode: String! ) {
        payOrder(orderID: $orderId, paymentMethodCode: $methodCode) {
            id
        }
    }
`);
