import {Slice, SliceCaseReducers, createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {AxiosError, AxiosResponse} from "axios";
import {createBackendClient} from "../../../common/backend/backendClient";
import {LoginRequest} from "../../../common/backend/requests/login-request";
import {LoginResponse} from "../../../common/backend/responses/login-response";
import {RefreshTokenResponse} from "../../../common/backend/responses/refresh-token-response";
import {jwt} from "../../../common/jwt/jwt";
import {AuthState, UserState} from "./state";


export const authLogin = createAsyncThunk('AUTH/LOGIN', async (loginRequest: LoginRequest, thunkAPI) => {
    return createBackendClient()
        .login(loginRequest)
        .then((res: AxiosResponse<LoginResponse>) => {
            return thunkAPI.fulfillWithValue(res.data);
        })
        .catch((err: AxiosError) => {
            return thunkAPI.rejectWithValue(err.response?.data);
        });
});

export const authRefreshToken = createAsyncThunk('AUTH/REFRESH-TOKEN', async (refreshToken: string, thunkAPI) => {
    return createBackendClient()
        .refreshToken(refreshToken)
        .then((res: AxiosResponse<RefreshTokenResponse>) => {
            return thunkAPI.fulfillWithValue(res.data);
        })
        .catch((err: AxiosError) => {
            return thunkAPI.rejectWithValue(err.response.data);
        });
});

const authSlice: Slice<AuthState> = createSlice<AuthState, SliceCaseReducers<AuthState>, string>({
    name: 'auth',
    initialState: {
        user: null,
        refreshToken: null,
        status: null,
        error: null
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(authLogin.pending, (state, action) => {
                state.status = 'pending';
                state.user = null;
            })
            .addCase(authLogin.rejected, (state, action) => {
                state.status = 'error';
                state.error = action.payload.status_code as string;
            })
            .addCase(authLogin.fulfilled, (state, action) => {
                const accessToken = action.payload.access_token;
                const refreshToken = action.payload.refresh_token;
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

                state.refreshToken = refreshToken;
                window.localStorage.setItem('refresh_token', refreshToken);

                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                // @ts-ignore
                state.user = decoded;
                state.status = 'fullfilled';
            })
            .addCase(authRefreshToken.pending, (state, action) => {
                state.status = 'pending';
                state.user = null;
            })
            .addCase(authRefreshToken.rejected, (state, action) => {
                state.status = 'error';
                state.error = action.payload.status_code as string;
            })
            .addCase(authRefreshToken.fulfilled, (state, action) => {
                const accessToken = action.payload.access_token;
                if (jwt.isExpired(accessToken)) {
                    state.status = 'error';
                    state.error = 'Token expired';
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

                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                // @ts-ignore
                state.user = decoded;
                state.status = 'fullfilled';
            });
    }
});

export default authSlice.reducer;