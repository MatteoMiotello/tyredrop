import {gql} from "@apollo/client";

export const USER_ADDRESS_FRAGMENT = gql`
    fragment UserAddressCollection on UserAddress {
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
`;