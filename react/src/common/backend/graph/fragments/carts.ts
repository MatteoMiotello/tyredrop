import {gql} from "../../../../__generated__";

export const CART_ITEMS_FRAGMENT = gql( /* GraphQL */ `
    fragment CartItems on Cart {
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
                brand {
                    name
                }
            }
        }
    }
`);