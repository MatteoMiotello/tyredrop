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
            priceAmount
            priceAmountTotal
            taxesAmount
            orderNumber
            addressName
            addressLine1
            addressLine2
            city
            country
            province
            postalCode
            status
            createdAt
            payment {
                id
                amount
                userPaymentMethod {
                    id
                    paymentMethod {
                        name
                        iban
                        bank_name
                        receiver
                    }
                }
            }
            userBilling {
                id
                name
                surname
                vatNumber
                fiscalCode
                legalEntityType {
                    name
                }
                user {
                    id
                    email
                }
            }
            orderRows {
                id
                amount
                quantity
                additionsAmount
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
    query fetchOrders( $userId: ID!, $pagination: PaginationInput, $filter: OrderFilterInput, $ordering: [OrderingInput] ) {
        userOrders(userId: $userId, pagination: $pagination, filter: $filter, ordering: $ordering) {
            data {
                id
                currency {
                    iso_code
                }
                status
                orderNumber
                addressName
                addressLine1    
                addressLine2
                city
                province
                postalCode
                country
                createdAt
                priceAmount
                taxesAmount
                priceAmountTotal
                orderRows {
                    id
                    amount
                    quantity
                }
            }
            pagination {
                ...PaginationInfo
            }
        }
    }
`;

export const FETCH_ALL_ORDERS = gql`
    ${PAGINATION_FRAGMENT}
    query allOrders( $filter: OrdersFilterInput, $pagination: PaginationInput, $ordering: [OrderingInput] ) {
        allOrders( filter: $filter, pagination: $pagination, ordering: $ordering ) {
            data {
                id
                priceAmount
                priceAmountTotal
                taxesAmount
                orderNumber
                createdAt
                status
                currency {
                    iso_code
                }
                userBilling {
                    name
                    surname
                    user {
                        id
                        email
                    }
                }
                orderRows {
                    id
                    amount
                }
            }
            pagination {
                ...PaginationInfo
            }
        }
    }
`;

export const POSSIBLE_ORDER_STATUSES = gql`
    query possibleOrderStatuses( $orderId: ID! ) {
        possibleOrderStatuses(orderId: $orderId)
    }
`;

export const ORDER_ROWS = gql`
    query orderRows( $orderId: ID! ) {
        orderRows( orderId: $orderId ) {
            id
            productItemPrice {
                id
                currency {
                    iso_code
                }
                value
                productItem {
                    id
                    product {
                        id
                        name
                        code
                    }
                    supplier {
                        id
                        name
                        code
                    }
                    supplierPrice
                }
            }
            amount
            additionsAmount
            quantity
        }
    }
`;