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
            productItemPrice {
                value
                currency {
                    iso_code
                }
                productItem {
                    id
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
    }
`;