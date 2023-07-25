import {gql} from "@apollo/client";

export const USER_CARTS = gql`
    query userCarts {
       carts {
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
                           brand {
                               name
                           }
                       }
                   }
               }
           }
       }
    }
`;