import React from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import {RegisterRequest} from "../../common/backend/requests/register-request";
import RegisterForm from "./components/RegisterForm";
import {useAuth} from "./hooks/useAuth";
import {authRegister} from "./store/auth-slice";
import {useToast} from "../../hooks/useToast";

const RegisterPage: React.FC = () => {
    const {t} = useTranslation();
    const dispatch = useDispatch<ThunkDispatch<RegisterRequest, any, any>>();
    const auth = useAuth();
    const navigate = useNavigate();
    const {setError} = useToast();

    
    const handleRegister = ( registerRequest: RegisterRequest ): void => {
        dispatch( authRegister( registerRequest ) );
    };

    const handleOnSuccess = () => {
        auth.tryRefreshToken();
    };

    return <>
        <div className="flex flex-col justify-center items-center my-auto">
            <h1 className="my-10">{t('register.page_title')}</h1>
            <RegisterForm
                register={handleRegister}
                onSuccess={handleOnSuccess}
            />
            <a className="link link-neutral link-hover text-secondary font-bold mt-5" href={"/auth/login"}> {t('register.login_label')} </a>
        </div>
    </>;
};

export default RegisterPage;