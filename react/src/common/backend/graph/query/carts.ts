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
       }
    }
`;