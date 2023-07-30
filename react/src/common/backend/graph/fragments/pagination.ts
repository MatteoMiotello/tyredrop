import {gql} from "@apollo/client";
export const PAGINATION_FRAGMENT = gql`
    fragment PaginationInfo on Pagination {
        limit
        totals
        offset
        currentPage
        pageCount
    }
`;