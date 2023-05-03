import {createSelector} from "@reduxjs/toolkit";
import {Store} from "../../../store/store";
import {Auth} from "../service/auth";
import {AuthState} from "./state";

const selectAuth = (state: Store) => state.auth;

const selectU = ( state: Store ) => state.auth.user;

export const selectUser = createSelector( [ selectAuth ], ( a: AuthState ) => {
    return a.user;
} );

export const selectAuthStatus = createSelector( [ selectAuth, selectU ], (a, u ): Auth => {
    return new Auth( a.status, a.error, a.refreshToken, u );
} );