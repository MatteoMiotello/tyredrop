
import {PAGINATION_FRAGMENT} from "../fragments/pagination";
import {gql} from "@apollo/client";

export const USER_ADDRESSES = gql`
    query fetchUserAddresses( $userId: ID! ) {
        userAddress( userId: $userId ) {
            ID
            addressName
            isDefault
            addressLine1
            addressLine2
            city
            country
            postalCode
            province
        }
    }
`;

export const USER_BILLING = gql`
    query fetchUserQuery( $userId: ID! ) {
        userBilling( userId: $userId ) {
            id
            name
            surname
            addressLine1
            addressLine2
            city
            country
            province
            cap
            fiscalCode
            vatNumber
            sdiCode
            sdiPec
            legalEntityType {
                name
            }
            user {
                id
                name
                surname
                email
            }
        }
    }
`;

export const USER = gql`
    query fetchUser( $userId: ID! ) {
        user( ID: $userId ) {
            id
            name
            surname
            email
            confirmed
            rejected
            avatarUrl
            userRole {
                isAdmin
                name
            }
            userBilling {
                id
                name
                surname
                addressLine1
                addressLine2
                city
                country
                province
                cap
                fiscalCode
                vatNumber
                sdiCode
                sdiPec
                legalEntityType {
                    name
                }
            }
        }
    }
`;

export const ALL_USERS = gql`
    ${PAGINATION_FRAGMENT}
    query fetchAllUsers( $pagination: PaginationInput, $filter: UserFilterInput ) {
        users( filter: $filter, pagination: $pagination) {
            data {
                id
                name
                surname
                email
                confirmed
                rejected
                userCode
                userRole {
                    id
                    roleCode
                    name
                }
            }
            pagination {
                ...PaginationInfo
            }
        }
    } 
`;

export const USER_BILLINGS = gql(/*GraphQL*/`
    query userBillings( $name: String ) {
        userBillings(name: $name) {
            id
            name
            surname
            vatNumber
            fiscalCode
        }
    } 
`);