import {
    AsyncThunk,
    PayloadAction,
    Slice,
    SliceCaseReducers,
    createAction,
    createAsyncThunk,
    createSlice
} from "@reduxjs/toolkit";
import {AxiosError, AxiosResponse} from "axios";
import {LoginRequest} from "../../../common/backend/requests/login-request";
import {RegisterRequest} from "../../../common/backend/requests/register-request";
import {LoginResponse} from "../../../common/backend/responses/login-response";
import {RefreshTokenResponse} from "../../../common/backend/responses/refresh-token-response";
import {UserStatus, extractFromJwt} from "../service/user";
import {AuthState} from "./state";
import {createBackendClient} from "../../../common/backend/backend-client";


export const authLogin: AsyncThunk<LoginResponse, any, any> = createAsyncThunk('AUTH/LOGIN', async (loginRequest: LoginRequest, thunkAPI) => {
    return createBackendClient()
        .login(loginRequest)
        .then((res: AxiosResponse<LoginResponse>) => {
            return thunkAPI.fulfillWithValue(res.data);
        })
        .catch((err: AxiosError) => {
            if (!err.response) {
                return thunkAPI.rejectWithValue(err.message);
            }

            return thunkAPI.rejectWithValue(err.response?.data);
        });
});

export const authRefreshToken: AsyncThunk<RefreshTokenResponse, any, any> = createAsyncThunk('AUTH/REFRESH-TOKEN', async (refreshToken: string, thunkAPI) => {
    return createBackendClient()
        .refreshToken(refreshToken)
        .then((res: AxiosResponse<RefreshTokenResponse>) => {
            return thunkAPI.fulfillWithValue(res.data);
        })
        .catch((err: AxiosError) => {
            if (!err.response) {
                return thunkAPI.rejectWithValue(err.message);
            }

            return thunkAPI.rejectWithValue(err.response.data);
        });
});

export const authRegister: AsyncThunk<LoginResponse, any, any> = createAsyncThunk('AUTH/REGISTER', async (request: RegisterRequest, thunkAPI) => {
    return createBackendClient()
        .signup(request)
        .then((res: AxiosResponse<LoginResponse>) => {
            return thunkAPI.fulfillWithValue(res.data);
        })
        .catch((err: AxiosError) => {
            if (!err.response) {
                return thunkAPI.rejectWithValue(err.message);
            }

            return thunkAPI.rejectWithValue(err.response.data);
        });
});

export const logout = createAction( 'AUTH/LOGOUT' );

export const setUserCompleted = createAction('AUTH/COMPLETE_USER');

const setupAuth = (state: AuthState, accessToken: string, refreshToken: string) => {
    let user = null;

    try {
        user = extractFromJwt(accessToken);
    } catch (error: any) {
        state.loggedIn = false;
        state.user = null;
        return;
    }

    state.refreshToken = refreshToken;
    window.localStorage.setItem('refresh_token', refreshToken);

    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    state.user = user;
    state.loggedIn = true;
    state.error = null;
};

const authSlice: Slice<AuthState> = createSlice<AuthState, SliceCaseReducers<AuthState>, string>({
    name: 'auth',
    initialState: {
        user: null,
        refreshToken: null,
        status: null,
        loggedIn: null,
        error: null
    },
    reducers: {},
    extraReducers: builder => {
        builder
            .addCase(authLogin.rejected, (state, action: PayloadAction<any>) => {
                state.loggedIn = false;
                if (action.payload && action.payload.status_code) {
                    state.error = action.payload.status_code as string;

                    return;
                }

                state.error = action.payload;
            })
            .addCase(authLogin.pending, (state) => {
                if ( state.error ) {
                    state.error = null;
                }
            })
            .addCase(authLogin.fulfilled, (state, action: PayloadAction<LoginResponse>) => {
                const accessToken = action.payload.access_token;
                const refreshToken = action.payload.refresh_token;

                setupAuth(state, accessToken, refreshToken);
            })
            .addCase(authRefreshToken.rejected, (state, action: PayloadAction<any>) => {
                state.loggedIn = false;
                state.error = action.payload.status_code as string;
            })
            .addCase(authRefreshToken.fulfilled, (state, action) => {
                const accessToken = action.payload.access_token;
                const refreshToken = action.payload.refresh_token;
                setupAuth(state, accessToken, refreshToken);
            })
            .addCase(authRegister.rejected, (state, action: PayloadAction<any>) => {
                state.loggedIn = false;
                state.error = action.payload.status_code as string;
            })
            .addCase(authRegister.pending, (state) => {
                if ( state.error ) {
                    state.error = null;
                }
            })
            .addCase(authRegister.fulfilled, (state, action: PayloadAction<LoginResponse>) => {
                const accessToken = action.payload.access_token;
                const refreshToken = action.payload.refresh_token;

                setupAuth(state, accessToken, refreshToken);
            })
            .addCase(setUserCompleted, (state) => {
                if (!state.user) {
                    return;
                }

                state.user.status = UserStatus.COMPLETED;
            })
            .addCase( logout, ( state ) => {
                window.localStorage.removeItem('refresh_token');
                state.loggedIn = false;
            } );
    }
});

export default authSlice.reducer;
