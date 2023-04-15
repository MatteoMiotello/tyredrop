import {HookHandler} from "vite";
import {useState} from "react";
import axios, {AxiosResponse} from "axios";
import backend from "../../config/backend";

interface RegisterRequest {
    email: string
    username: string
    password: string
    name: string
    surname: string
    language_code: string
}

interface RegisterResponse {
    email: string
    username: string
    name: string
    surname: string
    role: string
}

const useRegister: HookHandler<any> = () => {
    const [data, setData] = useState<RegisterResponse>();

    const path = '/register';

    const handleRegister = (payload: RegisterRequest) => {
        axios.post(backend.endpoint + path, payload)
            .then((res: AxiosResponse<RegisterResponse>) => setData(res.data));
    };

    return [data, handleRegister];
};