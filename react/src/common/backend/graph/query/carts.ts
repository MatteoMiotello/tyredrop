import {gql} from "../../../../__generated__";

export const USER_CARTS = gql(`
query userCarts {
   carts {
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