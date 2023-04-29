import {createSelector} from "@reduxjs/toolkit";
import {Store} from "../../../store/store";
import {AuthStatus} from "../service/auth-status";
import {AuthState} from "./state";

const selectAuth = (state: Store) => state.auth;

export const selectUser = createSelector( [ selectAuth ], ( a: AuthState ) => {
    return a.user;
} );

export const selectAuthStatus = createSelector( [ selectAuth ], (a: AuthState ): AuthStatus => {
    return new AuthStatus( a.status, a.error );
} );