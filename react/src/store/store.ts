import {configureStore} from "@reduxjs/toolkit";
import {createBackendClient} from "../common/backend/backendClient";
import authSlice from "../modules/auth/store/auth-slice";
import {AuthState} from "../modules/auth/store/state";

export type Store = {
    auth: AuthState
}
export const store = configureStore({
    reducer: {
        auth: authSlice
    },
    middleware: getDefaultMiddleware => getDefaultMiddleware({
        thunk: {
            extraArgument: {
                backend: createBackendClient()
            }
        }
    })
});