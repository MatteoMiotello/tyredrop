import {gql} from "../../../../__generated__";

export const UPDATE_MARKUP = gql(/* GraphQL */`
    mutation updateMarkup( $id: ID!, $markupPercentage: Int! ) {
        updatePriceMarkup(id: $id, markupPercentage: $markupPercentage) {
            id
        }   
    }
`);