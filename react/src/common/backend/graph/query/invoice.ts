import {gql} from "@apollo/client";


export const ALL_INVOICES = gql( /* GraphQL */`
    query allInvoices( $pagination: PaginationInput!, $input: InvoiceFilter! ) {
        allInvoices(input: $input, pagination: $pagination) {
            data {
                id
                number
                fileUrl
                createdAt
                userBilling {
                    id
                    user {
                        id
                        email
                    }
                    name
                    surname
                }
            }
            pagination {
                limit
                offset
                pageCount
                currentPage
                totals
            }
        }
    }
`);

export const ALL_USER_INVOICES = gql(/* GraphQL */`
    query allUserInvoices( $pagination: PaginationInput!, $input: InvoiceFilter! ) {
        allInvoices(input: $input, pagination: $pagination) {
            data {
                id
                number
                fileUrl
                createdAt
            }
            pagination {
                limit
                offset
                pageCount
                currentPage
                totals
            }
        }
    }
`);