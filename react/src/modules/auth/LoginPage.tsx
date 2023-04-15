import React from "react";
import Panel from "../../common/components-library/Panel";
import Input, {ValidateHandler} from "../../common/components-library/Input";
import {useTranslation} from "react-i18next";
import Button from "../../common/components-library/Button";
import logo from "../../assets/logo-transparent.png";
import useLogin from "../../hooks/backend/useLogin";
import Form, {useForm} from "../../common/components-library/Form";

export const LoginPage: React.FC = () => {
    const [data, handleLogin] = useLogin();
    const [form, handleFormError] = useForm();
    const {t} = useTranslation();

    const validateUsername: ValidateHandler = ( value ) => {
        if ( !value ) {
            return t( 'login.username_required' );
        }

        return null;
    };

    function onSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);

        return;

        const username = formData.get("username");
        const password = formData.get("password");

        handleLogin(username, password);
    }

    if (data) {
        console.log(data);
    }

    return <div>
        <main className="bg-base lg:p-24 p-4">
            <Panel className="flex flex-col justify-center items-center my-auto">
                <img src={logo} width={100}/>
                <Form onSubmit={onSubmit} form={form}>
                    <Input name="username"
                           type="text"
                           placeholder={t('login.username_placeholder')}
                           required={true}
                           className="col-span-12 lg:col-start-4 lg:col-span-6"
                           validate={ validateUsername }
                    />
                    <Input name="password"
                           type="password"
                           placeholder="Password"
                           required={true}
                           className="col-span-12 lg:col-start-4 lg:col-span-6"
                    />
                    <div className="flex justify-between w-full text-sm col-span-12 lg:col-start-4 lg:col-span-6">
                        <a className="link link-neutral link-hover"
                           href={"/auth/register"}> {t('login.register_label')} </a>
                        <a className="link link-neutral link-hover"> {t('login.forgot_password')} </a>
                    </div>
                    <Button
                        type={"primary"}
                        htmlType={"submit"}
                        className="col-span-12 lg:col-start-4 lg:col-span-6"
                    >
                        {t('login.Submit')}
                    </Button>
                </Form>
            </Panel>
        </main>
    </div>;
};