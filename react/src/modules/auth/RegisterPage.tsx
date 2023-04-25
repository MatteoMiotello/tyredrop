import React from "react";
import {useTranslation} from "react-i18next";
import RegisterForm from "./components/RegisterForm";

const RegisterPage: React.FC = () => {
    const {t} = useTranslation();

    return <>
        <div className="flex flex-col justify-center items-center my-auto">
            <h1 className="my-10">{t('register.page_title')}</h1>
            <RegisterForm/>
            <a className="link link-neutral link-hover text-secondary font-bold mt-5" href={"/auth/login"}> {t('register.login_label')} </a>
        </div>
    </>;
};

export default RegisterPage;