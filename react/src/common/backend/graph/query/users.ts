
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