import {gql} from "../../../../__generated__";
import {CART_ITEMS_FRAGMENT} from "../fragments/carts";

export const USER_CARTS = gql( /* GraphQL */ `
    query userCarts {
       carts {
           ...CartItems
        }
    }
    ${CART_ITEMS_FRAGMENT}
`);