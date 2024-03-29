import {gql} from "../../../../__generated__";

export const CREATE_INVOICE = gql( /* GraphQL */ `
    mutation createInvoice( $billingId: ID!, $number: String, $file: Upload! ) {
        createInvoice( userBillingId: $billingId, number: $number, file: $file ) {
            id
        }
    }
`);

export const DELETE_INVOICE = gql( /*GraphQL*/`
    mutation deleteInvoice( $id: ID! ) {
        deleteInvoice(id: $id) {
            id
        }
    }
` );

export const UPDATE_INVOICE_STATUS = gql( /*GraphQL*/ `
    mutation updateInvoiceStatus( $id: ID!, $status: InvoiceStatus! ) {
        updateInvoiceStatus(id: $id, status: $status) {
            id
        }
    }
`);