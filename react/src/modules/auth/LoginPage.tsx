import {AnyAction} from "@reduxjs/toolkit";
import React from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import {LoginRequest} from "../../common/backend/requests/login-request";
import {Store} from "../../store/store";
import LoginForm from "./components/LoginForm";
import {authLogin} from "./store/auth-slice";

const LoginPage: React.FC = () => {
    const dispatch: ThunkDispatch<Store, any, AnyAction> = useDispatch();
    const {t} = useTranslation();
    
    const login = ( loginRequest: LoginRequest ) => {
        dispatch(authLogin( loginRequest ));
    };

    return <>
            <div className="flex flex-col justify-center items-center my-auto">
                <h1 className="my-10">{t('login.page_title')}</h1>
                <LoginForm login={login}/>
                <a className="link link-neutral link-hover text-secondary font-bold mt-5" href={"/auth/register"}> {t('login.register_label')} </a>
            </div>
    </>;
};

export default LoginPage;