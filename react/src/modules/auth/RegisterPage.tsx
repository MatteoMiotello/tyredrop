import React, {useEffect} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch, useSelector} from "react-redux";
import {useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import {RegisterRequest} from "../../common/backend/requests/register-request";
import Spinner from "../../common/components/Spinner";
import RegisterForm from "./components/RegisterForm";
import {useAuthenticated} from "./hooks/useAuthenticated";
import {selectAuthStatus} from "./store/auth-selector";
import {authRegister} from "./store/auth-slice";

const RegisterPage: React.FC = () => {
    const {t} = useTranslation();
    const dispatch = useDispatch<ThunkDispatch<RegisterRequest, any, any>>();
    const isAuthenticated = useAuthenticated();
    const authStatus = useSelector( selectAuthStatus );
    const navigate = useNavigate();

    useEffect( () => {
        if ( isAuthenticated ) {
            navigate( '/' );
        }
    }, []);

    const handleRegister = ( registerRequest: RegisterRequest ): void => {
        dispatch( authRegister( registerRequest ) );
    };

    const handleOnSuccess = () => {
        navigate( '/' );
    };

    return <>
        <div className="flex flex-col justify-center items-center my-auto">
            <h1 className="my-10">{t('register.page_title')}</h1>
            { authStatus.status == 'pending' && <Spinner/> }
            <RegisterForm
                register={handleRegister}
                onSuccess={handleOnSuccess}
            />
            <a className="link link-neutral link-hover text-secondary font-bold mt-5" href={"/auth/login"}> {t('register.login_label')} </a>
        </div>
    </>;
};

export default RegisterPage;