import {gql} from "../../../../__generated__";

export const GET_LEGAL_ENTITY_TYPES = gql(`
    query GetLegalEntityTypes {
        legalEntityTypes {
            id
            name
        }
    }
`);