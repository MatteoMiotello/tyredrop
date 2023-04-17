import {configureStore} from "@reduxjs/toolkit";
import {createBackendClient} from "../common/backend/backendClient";
import authSlice, {AuthState} from "../modules/auth/store/authSlice";

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