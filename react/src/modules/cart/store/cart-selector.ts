import {createSelector} from "@reduxjs/toolkit";
import {Store} from "../../../store/store";

const cartItems = ( store: Store ) => store.cart.items;

const selectCartCount = createSelector( [ cartItems ], ( carts ) => {
    if ( !carts ) {
        return 0;
    }

    return carts.length;
} );

const selectCarts = createSelector( [ cartItems ], ( carts ) => {
    return carts;
} );

const cartSelectors = {
    count: selectCartCount,
    items: selectCarts
};

export default cartSelectors;

