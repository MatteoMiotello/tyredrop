import React, {useState} from "react";
import {HookHandler} from "vite";
import axios, {AxiosResponse} from "axios";
import backend from "../../config/backend";
import {use} from "i18next";

interface LoginResponse {
    access_token: string,
    refresh_token: string
}

const useLogin: HookHandler<any> = () => {
    const [data, setData] = useState<LoginResponse|null>(null);
    const path = '/auth';

    const handleLogin = function ( username: string, password: string ) {
        axios.post(backend.endpoint + path, {
            username: username,
            password: password
        }).then( ( res: AxiosResponse<LoginResponse> ) => {
            setData( res.data);
        } );
    };

    return [data, handleLogin];
};

export default useLogin;