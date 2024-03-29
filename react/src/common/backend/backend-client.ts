import axios, {AxiosInstance, AxiosResponse} from "axios";
import {LoginRequest} from "./requests/login-request";
import {RegisterRequest} from "./requests/register-request";
import {LoginResponse} from "./responses/login-response";
import {RefreshTokenResponse} from "./responses/refresh-token-response";
import config from "../../config";
import {SupportRequest} from "./requests/support-request";

export interface BackendClient {
    login(request: LoginRequest): Promise<AxiosResponse<LoginResponse>>;
    refreshToken(refreshToken: string): Promise<AxiosResponse<RefreshTokenResponse>>
    signup( request: RegisterRequest ): Promise<AxiosResponse<LoginResponse>>
}

export const createBackendClient = () => {
    return new Backend();
};

class Backend implements BackendClient {
    private client: AxiosInstance;

    constructor() {
        this.client = axios.create({
            baseURL: config.backend.endpoint,
            withCredentials: true
        });
    }

    private makePostRequest<T = any>(path: string, body: any): Promise<AxiosResponse<T>> {
        return this.client.post(path, body);
    }

    login(request: LoginRequest): Promise<AxiosResponse<LoginResponse>> {
        return this.makePostRequest<LoginResponse>('/login', request);
    }

    signup( request: RegisterRequest ): Promise<AxiosResponse<LoginResponse>> {
        return this.makePostRequest<LoginResponse>( '/register', request );
    }

    refreshToken(refreshToken: string): Promise<AxiosResponse<RefreshTokenResponse>> {
        return this.makePostRequest<RefreshTokenResponse>( '/refresh_token', {refresh_token: refreshToken} );
    }

    resetPassword( email: string ): Promise<AxiosResponse> {
        return this.makePostRequest( '/reset_password', {email: email} );
    }

    changePassword( token: string, password: string ): Promise<AxiosResponse> {
        return this.makePostRequest( '/change_password', {token: token, password: password} );
    }

    supportEmail( request: SupportRequest ): Promise<AxiosResponse<void>> {
        return this.makePostRequest( '/support_request', request );
    }
}