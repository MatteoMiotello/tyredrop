
import {gql} from "@apollo/client";
import {USER_ADDRESS_FRAGMENT} from "../fragments/users";

export const USER_ADDRESSES = gql`
    query fetchUserAddresses {
        userAddress {
            ...UserAddressCollection
        }
    }
    ${USER_ADDRESS_FRAGMENT}
`;

export const USER_BILLING = gql`
    query fetchUserQuery {
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
            legalEntityType {
                name
            }s
        }
    }
`;