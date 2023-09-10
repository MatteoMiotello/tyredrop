import {gql} from "../../../../__generated__";

export const UPDATE_MARKUP = gql(/* GraphQL */`
    mutation updateMarkup( $id: ID!, $input: PriceMarkupInput! ) {
        updatePriceMarkup(id: $id, input: $input) {
            id
        }   
    }
`);

export const CREATE_MARKUP = gql( /* GraphQL */`
    mutation createMarkup( $input: PriceMarkupInput! ) {
        createPriceMarkup(input: $input){
            id
        }
    }
` );