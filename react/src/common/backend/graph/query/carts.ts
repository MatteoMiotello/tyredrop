import {gql} from "@apollo/client";
import {CART_ITEMS_FRAGMENT} from "../fragments/carts";

export const USER_CARTS = gql`
    ${CART_ITEMS_FRAGMENT}
    query userCarts {
       carts {
           ...CartItems
       }
    }
`;