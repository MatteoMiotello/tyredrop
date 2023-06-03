import {createSelector} from "@reduxjs/toolkit";
import {Store} from "../../../store/store";

const cartItems = ( store: Store ) => store.cart.items;
const amountTotal = ( store: Store ) => store.cart.amountTotal;

const selectCartCount = createSelector( [ cartItems ], ( carts ) => {
    if ( !carts ) {
        return 0;
    }

    return carts.length;
} );

const selectCarts = createSelector( [ cartItems ], ( carts ) => {
    return carts;
} );

const selectCart = ( itemId: string ) => createSelector( [ cartItems ], ( carts ) => {
    if ( !carts.length ) {
        return null;
    }

    return carts.find( ( cart ) => {return cart.id === itemId;} );
} );

const selectAmountTotal = createSelector( [amountTotal], ( amountTotal ) => {
    return amountTotal;
} );

const cartSelectors = {
    count: selectCartCount,
    items: selectCarts,
    cart: selectCart,
    amount: selectAmountTotal
};

export default cartSelectors;

