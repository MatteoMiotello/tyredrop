import {gql} from "../../../../__generated__";

export const ADD_CART = gql(`
    mutation addCart( $itemId: ID!, $quantity: Int) {
        addItemToCart( itemId: $itemId, quantity: $quantity ) {
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
                }
            }
        } 
    } 
`);