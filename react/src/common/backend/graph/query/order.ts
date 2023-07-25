import {PAGINATION_FRAGMENT} from "../fragments/pagination";
import {gql} from "@apollo/client";


export const FETCH_ORDER = gql `
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
            createdAt
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
                productItemPrice {
                    value
                    currency {
                        iso_code
                        symbol
                    }
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
    }
`;

export const FETCH_USER_ORDERS = gql`
    ${PAGINATION_FRAGMENT}
    query fetchOrders( $userId: ID!, $pagination: PaginationInput ) {
        userOrders(userId: $userId, pagination: $pagination) {
            data {
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
                createdAt
                orderRows {
                   amount
                }
            }
            pagination {
                ...PaginationInfo
            }
        }
    }
`;