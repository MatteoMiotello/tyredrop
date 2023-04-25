import {createSelector} from "@reduxjs/toolkit";
import {Store} from "../../../store/store";
import {AuthState} from "./state";

export interface UserStatus {
    error: string | null | number,
    status: string | null | undefined
}

const selectAuth = (state: Store) => state.auth;

export const selectUser = createSelector( [ selectAuth ], ( a: AuthState ) => {
    return a.user;
} );

export const selectUserStatus = createSelector( [ selectAuth ], ( a: AuthState ): UserStatus => {
    return {status: a.status, error: a.error};
} );