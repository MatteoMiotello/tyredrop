import {gql} from "@apollo/client";
import {CART_ITEMS_FRAGMENT} from "../fragments/carts";

export const ADD_CART = gql`
    ${CART_ITEMS_FRAGMENT} 
    mutation addCart( $itemId: ID!, $quantity: Int) {
        addItemToCart( itemId: $itemId, quantity: $quantity ) {
           ...CartItems
        }
    } 
`;

export const EDIT_CART = gql`
    ${CART_ITEMS_FRAGMENT}
    mutation editCart( $cartId: ID!, $quantity: Int! ) {
        editCart( cartId: $cartId, quantity: $quantity ) {
            ...CartItems
        }
    } 
`;