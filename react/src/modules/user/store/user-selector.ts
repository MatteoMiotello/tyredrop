import {createSelector} from "@reduxjs/toolkit";
import {Store} from "../../../store/store";

const selecUser = (state: Store) => state.user;

const selectUserAddresses = createSelector([selecUser], (user) => {
    return user.addresses;
});

const userSelectors = {
    addresses: selectUserAddresses
};

export default userSelectors;