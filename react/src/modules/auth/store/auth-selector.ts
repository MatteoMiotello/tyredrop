import {createSelector} from "@reduxjs/toolkit";
import {Store} from "../../../store/store";
import {AuthState} from "./state";

export interface AuthStatus {
    error: string | null | number,
    status: string | null | undefined
}

const selectAuth = (state: Store) => state.auth;

export const selectUser = createSelector( [ selectAuth ], ( a: AuthState ) => {
    return a.user;
} );

export const selectAuthStatus = createSelector( [ selectAuth ], (a: AuthState ): AuthStatus => {
    return {status: a.status, error: a.error};
} );