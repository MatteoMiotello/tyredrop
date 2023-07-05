import {gql} from "../../../../__generated__";

export const FETCH_ORDER = gql(`
    query fetchOrder( $orderId: ID! ) {
        order(id: $orderId) {
            id
            currency {
                iso_code
                name
            }
            addressLine1
            addressLine2
            city
            country
            province
            postalCode
            status
            userBilling {
                id
                name
                surname
                vatNumber
                fiscalCode
                legalEntityType {
                    name
                }
            }
            orderRows {
                id
                amount
                quantity
                productItem {
                    id
                    product {
                        id
                        name
                    }
                }
            }
        }
    }
`);

export const FETCH_USER_ORDERS = gql(`
    query fetchOrders( $userId: ID!, $pagination: PaginationInput ) {
        userOrders(userId: $userId, pagination: $pagination) {
            id
            currency {
                iso_code
            }
            status
            addressLine1
            addressLine2
            city
            province
            postalCode
            country
            orderRows {
               amount
            }
        }
    }
`);