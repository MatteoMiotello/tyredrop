import {gql} from "../../../../__generated__";


export const ALL_PAYMENT_METHODS = gql(`
    query allMethods {
        paymentMethods {
            id
            code
            name
            receiver
            bank_name
            iban
        }
    }
`);