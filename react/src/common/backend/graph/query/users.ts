
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
                legalEntityType {
                    name
                }
            }
        }
    }
`;