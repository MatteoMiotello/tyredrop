import React, {useEffect} from "react";
import {LoginRequest} from "../../../common/backend/requests/login-request";
import Input, {ValidationHandler} from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import Form, {FormErrors, FormSubmitHandler, useForm} from "../../../common/components-library/Form";
import {useTranslation} from "react-i18next";
import {useSelector} from "react-redux";
import {Store} from "../../../store/store";
import {selectUserStatus} from "../store/auth-selector";
import Spinner from "../../../common/components/Spinner";

interface LoginFormProps {
    login: ( request: LoginRequest ) => void

    onSuccess: () => void
}

const LoginForm: React.FC<LoginFormProps> = ( props: LoginFormProps ) => {
    const [form, handleFormError] = useForm();
    const {t} = useTranslation();
    const userStatus = useSelector<Store>( selectUserStatus );

    useEffect( () => {
        if ( userStatus.status == 'error' ) {
            let error = userStatus.error;
            if ( userStatus.error >= 4000 ) {
                error = t( 'login.wrong_username_or_password' );
            }

            handleFormError( error );
        }

        if ( userStatus.status == 'fullfilled' ) {
            props.onSuccess();
        }
    }, [ userStatus ] );

    const validateUsername: ValidationHandler = ( value ) => {
        if ( !value ) {
            return t( 'login.username_required' );
        }

        return null;
    };

    const validatePassword: ValidationHandler = ( value ) => {
        if ( !value ) {
            return t( 'login.password_required' );
        }

        return null;
    };

    const onSubmit: FormSubmitHandler<LoginRequest> = ( loginRequest: LoginRequest ) => {
        const formErrors = new FormErrors();

        formErrors.appendError( validateUsername( loginRequest.username as string ) );
        formErrors.appendError( validatePassword( loginRequest.password as string ) );

        handleFormError( formErrors );

        if ( formErrors.hasErrors() ) {
            return;
        }

        props.login( loginRequest );
    };

    return <Form onSubmit={(r) => onSubmit(r)} form={form} className={"relative"}>
        { (userStatus.status == 'pending') && <Spinner></Spinner> }
        <Input name="username"
               type="text"
               placeholder={t('login.username_placeholder')}
               className="col-span-12 lg:col-start-4 lg:col-span-6"
               validate={ validateUsername }
        />
        <Input name="password"
               type="password"
               placeholder="Password"
               className="col-span-12 lg:col-start-4 lg:col-span-6"
               validate={ validatePassword }
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
    </Form>;
};

export default LoginForm;