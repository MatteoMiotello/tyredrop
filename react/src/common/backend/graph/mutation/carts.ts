import {gql} from "../../../../__generated__";
import {CART_ITEMS_FRAGMENT} from "../fragments/carts";

export const ADD_CART = gql( /* GraphQL */  `
    mutation addCart( $itemId: ID!, $quantity: Int) {
        addItemToCart( itemId: $itemId, quantity: $quantity ) {
           ...CartItems
        }
    } 
    ${CART_ITEMS_FRAGMENT} 
`);