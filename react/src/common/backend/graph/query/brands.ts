import {gql} from "../../../../__generated__";

export const SEARCH_BRANDS = gql(`
     query SearchBrands( $name: String! ) {
        searchBrands(  name: $name ) {
            id
            name
            code
            quality
        }
     }  
`);