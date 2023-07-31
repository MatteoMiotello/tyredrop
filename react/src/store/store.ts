import {configureStore} from "@reduxjs/toolkit";
import authSlice from "../modules/auth/store/auth-slice";
import {AuthState} from "../modules/auth/store/state";
import {createBackendClient} from "../common/backend/backend-client";
import UserState from "../modules/user/store/state";
import userSlice from "../modules/user/store/user-slice";
import appSlice from "./app-slice";
import {ProductCategory} from "../__generated__/graphql";
import {CartState} from "../modules/cart/store/state";
import cartSlice from "../modules/cart/store/cart-slice";
import {ToastState, toastMiddleware} from "./toast";
import toastSlice from "./toast/toast-slice";


export type AppState = {
    productCategories: ProductCategory[]
}

export type Store = {
    auth: AuthState
    app: AppState
    cart: CartState
    user: UserState
    toast: ToastState
}
export const store = configureStore({
    reducer: {
        auth: authSlice,
        app: appSlice,
        cart: cartSlice,
        user: userSlice,
        toast: toastSlice,
    },
    middleware: getDefaultMiddleware => getDefaultMiddleware({
        thunk: {
            extraArgument: {
                backend: createBackendClient()
            }
        },
    }).concat(toastMiddleware)
});