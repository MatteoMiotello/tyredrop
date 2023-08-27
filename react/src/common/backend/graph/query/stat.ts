import {gql} from "../../../../__generated__";

export const STATS = gql(/* GraphQL */`
    query stats {
        stats {
            totalOrders
            totalUsers
            bestUser {
                id
                name
                surname
                user {
                    id
                    email
                }
            }
        }
    }
`);