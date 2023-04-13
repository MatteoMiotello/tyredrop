import React, {FormEvent, FormEventHandler} from "react";
import Panel from "../../common/components/Panel";
import Input from "../../common/components/Input";
import {useTranslation} from "react-i18next";
import Button from "../../common/components/Button";
import logo from "../../assets/logo-transparent.png"
import Footer from "../../common/components/Footer";
import CustomFooter from "../../common/components/CustomFooter";
import useLogin from "../../hooks/backend/useLogin";
import {useNavigate} from "react-router-dom";

export const LoginPage: React.FC = () => {
    const [ data, handleLogin ] = useLogin()
    const {t, i18n} = useTranslation()
    let navigate = useNavigate()

    function onSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()
        const formData = new FormData(e.currentTarget);

        const username = formData.get( "username" )
        const password = formData.get( "password" )

        handleLogin( username, password )
    }

    if ( data ) {
        console.log( data )
    }

    return <div>
        <main className="bg-base md:p-24 p-4 h-screen">
            <Panel className="flex flex-col justify-center items-center my-auto">
                <img src={logo} width={100}/>
                <form className={"flex flex-col justify-center items-center md:w-1/2 xs:w-full"} onSubmit={onSubmit}>
                    <Input name="username" type="text" placeholder={t('login.username_placeholder')}
                           required={true}></Input>
                    <Input name="password" type="password" placeholder="Password" required={true}></Input>
                    <div className="flex justify-between w-full text-sm mx-2">
                        <a className="link link-neutral link-hover" href={"/auth/register"}> {t('login.register_label')} </a>
                        <a className="link link-neutral link-hover"> {t('login.forgot_password')} </a>
                    </div>
                    <Button type={"primary"} htmlType={"submit"}>
                        {t('login.Submit')}
                    </Button>
                </form>
            </Panel>
        </main>
    </div>
}