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

export const EDIT_USER_ADDRESS  = gql`
    ${USER_ADDRESS_FRAGMENT}
    mutation editAddress( $id: ID!, $input: UserAddressInput! ) {
        editUserAddress(id: $id, userAddress: $input) {
            ...UserAddressCollection
        }
    }
`;

export const DELETE_USER_ADDRESS = gql`
    ${USER_ADDRESS_FRAGMENT}
    mutation deleteAddress( $id: ID! ) {
        deleteUserAddress(id: $id) {
            ...UserAddressCollection
        }
    }
`;