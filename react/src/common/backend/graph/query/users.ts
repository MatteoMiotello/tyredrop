import {gql} from "@apollo/client";

export const USER_ADDRESSES = gql`
    query fetchUserAddresses {
        userAddress {
            ID
            isDefault
            addressLine1
            addressLine2
            country
            city
            postalCode
            province
        }
    }
`;