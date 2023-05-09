import {configureStore} from "@reduxjs/toolkit";
import authSlice from "../modules/auth/store/auth-slice";
import {AuthState} from "../modules/auth/store/state";
import {createBackendClient} from "../common/backend/backend-client";
import appSlice from "./app-slice";
import {ProductCategory} from "../__generated__/graphql";


export type AppState =  {
    productCategories: ProductCategory[]
}

export type Store = {
    auth: AuthState
    app: AppState
}
export const store = configureStore({
    reducer: {
        auth: authSlice,
        app: appSlice
    },
    middleware: getDefaultMiddleware => getDefaultMiddleware({
        thunk: {
            extraArgument: {
                backend: createBackendClient()
            }
        }
    })
});