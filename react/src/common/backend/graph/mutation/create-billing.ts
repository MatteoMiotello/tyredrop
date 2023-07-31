import {gql} from "../../../../__generated__";

export const CREATE_BILLING = gql(`
    mutation CreateUserBilling( $input: CreateUserBilling! ) {
        createUserBilling( billingInput: $input ) {
            id
            name
            surname
        }
    }
`);