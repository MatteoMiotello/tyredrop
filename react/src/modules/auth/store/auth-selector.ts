import {createSelector} from "@reduxjs/toolkit";
import {Store} from "../../../store/store";
import {AuthState} from "./state";

const selectAuth = (state: Store) => state.auth;

export const selectUser = createSelector( [ selectAuth ], ( a: AuthState ) => {
    return a.user;
} );

export const selectUserStatus = createSelector( [ selectAuth ], ( a: AuthState ) => {
    return {status: a.status, error: a.error};
} );