import {gql} from "@apollo/client";


export const CART_ITEMS_FRAGMENT = gql`
    fragment CartItems on CartResponse {
        totalPrice {
            value
            currency {
                iso_code
                symbol
                name
            }
        }
        items {
            id
            quantity
            productItem {
                id
                price {
                    value
                    currency {
                        iso_code
                    }
                }
                product {
                    name
                    code
                    eprelProductCode
                    brand {
                        name
                    }
                }
            }
        }
    }
`;