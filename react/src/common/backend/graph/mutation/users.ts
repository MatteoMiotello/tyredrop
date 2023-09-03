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

export const UPDATE_USER_STATUS = gql`
    mutation changeUserStatus( $userID: ID!, $confirmed: Boolean, $rejected: Boolean ) {
        updateUserStatus( userID: $userID, confirmed: $confirmed, rejected: $rejected ) {
            id
        }
    }
`;

export const UPDATE_AVATAR = gql`
    mutation updateAvatar( $userID: ID!, $file: Upload! ) {
        updateAvatar(userID: $userID, file: $file) {
            id
        }
    }
`;

export const UPDATE_USER_BILLING = gql`
    mutation updateUserBilling( $billingID: ID!, $input: BillingInput ) {
        updateUserBilling(userBillingID: $billingID, billingInput: $input) {
            id
        }
    }
`;