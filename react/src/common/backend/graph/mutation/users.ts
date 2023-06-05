import {gql} from "@apollo/client";
import {USER_ADDRESS_FRAGMENT} from "../fragments/users";

export const ADD_USER_ADDRESS = gql`
    ${USER_ADDRESS_FRAGMENT}
    mutation addAddress( $input: UserAddressInput! ) {
        createUserAddress(userAddress: $input) {
            ...UserAddressCollection
        }
    }
`;