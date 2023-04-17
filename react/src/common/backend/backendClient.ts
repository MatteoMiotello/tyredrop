import axios, {AxiosInstance, AxiosResponse} from "axios";
import backend from "../../config/backend";
import {LoginRequest} from "./requests/LoginRequest";
import {LoginResponse} from "./responses/LoginResponse";

export const createBackendClient = () => {
    return new BackendClient();
};

class BackendClient {
    private client: AxiosInstance;

    constructor() {
        this.client = axios.create({
            baseURL: backend.endpoint
        });
    }

    private makePostRequest<T = any>(path: string, body: any): Promise<AxiosResponse<T>> {
        return this.client.post(path, body);
    }

    login(request: LoginRequest): Promise<AxiosResponse<LoginResponse>> {
        return this.makePostRequest<LoginResponse>('/login', request);
    }
}