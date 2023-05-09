import {gql} from "../../../../__generated__";

export const ALL_CATEGORIES_WITH_SPECIFICATIONS = gql(`
    query getAllCategories {
        productCategories {
            id
            name
            code
            specifications {
                id
                code
                name
                type
                mandatory
                searchable
            }
        }
    }
`);