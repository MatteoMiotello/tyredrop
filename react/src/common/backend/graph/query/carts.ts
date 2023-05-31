import {gql} from "../../../../__generated__";

export const USER_CARTS = gql(`
    query userCarts() {
        carts {
           productItem {
            id
            product {
                id
                code
                name
                brand {
                    name
                    code
                }
            }
        }
    }
`);

export const ADD_CART = gql(`
    mutation addItemToCart( $itemId: ID!, $quantity: Int){
        addItemToCart( itemId: $itemId, $quantity: $quantity ) {
            id
            quantity
            productItem {
                id
            }
        } 
    } 
`);