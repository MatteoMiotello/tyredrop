import {Slice, SliceCaseReducers, createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {createBackendClient} from "../../../common/backend/backendClient";
import {LoginRequest} from "../../../common/backend/requests/LoginRequest";
import {jwt} from "../../../common/jwt/jwt";


export type AuthState = {
    user: UserState | null
    status: string
}
export type UserState = {
    username: string
    email: string
    name?: string
    surname?: string
    role: string
}

export const authLogin = createAsyncThunk('auth/login', async (loginRequest: LoginRequest) => {
    const res = await createBackendClient().login(loginRequest);

    return res.data;
});
const authSlice: Slice<AuthState> = createSlice<AuthState, SliceCaseReducers<AuthState>, string>({
    name: 'auth',
    initialState: {
        user: null,
        status: 'pending'
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(authLogin.pending, (state, action) => {
                state.status = 'pending';
                state.user = null;
            })
            .addCase(authLogin.rejected, (state, action) => {
                state.status = action.error.message as string;
            })
            .addCase(authLogin.fulfilled, (state, action) => {
                const accessToken = action.payload.access_token;

                if (jwt.isExpired(accessToken)) {
                    state.status = 'Token expired';
                    state.user = null;
                    return;
                }
                let decoded = null;
                try {
                    decoded = jwt.decodeJwt<UserState>(accessToken);
                } catch (error: any) {
                    state.status = error.message;
                    return;
                }

                state.user = decoded;
            });
    }
});

export default authSlice.reducer;