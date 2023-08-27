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

export const UPDATE_ORDER_STATUS = gql( /* GraphQL */ `
    mutation updateOrderStatus( $orderId: ID!, $newStatus: OrderStatus! ) {
        updateOrderStatus( orderID: $orderId, newStatus: $newStatus ) {
            id
        }
    }
` );

export const UPDATE_ORDER_ROW = gql( /* GraphQL */`
    mutation updateOrderRow( $rowId: ID!, $input: OrderRowInput! ) {
        updateOrderRow(input: $input, rowID: $rowId) {
            id
        }
    }
` );