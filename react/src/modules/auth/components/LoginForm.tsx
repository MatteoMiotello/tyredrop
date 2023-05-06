import React, { useEffect} from "react";
import {LoginRequest} from "../../../common/backend/requests/login-request";
import Field from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import Form, {FormErrors, FormSubmitHandler, useForm} from "../../../common/components-library/Form";
import {useTranslation} from "react-i18next";
import {ValidationHandler} from "../../../common/validation/validators";
import {useAuth} from "../hooks/useAuth";

interface LoginFormProps {
    login: ( request: LoginRequest ) => void

    onSuccess?: () => void
}

const LoginForm: React.FC<LoginFormProps> = ( props: LoginFormProps ) => {
    const [form, handleFormError] = useForm();
    const {t} = useTranslation();
    const auth = useAuth();

    useEffect( () => {
        if ( auth.isError() ) {
            let error = auth.error;
            if ( auth.error && ( typeof auth.error == 'number' && auth.error >= 4000 ) ) {
                error = t( 'login.wrong_username_or_password' );
            }

            handleFormError( error );
        }

        if ( auth.isAuthenticated() && props.onSuccess) {
            props.onSuccess();
        }
    }, [ auth ] );

    const validateEmail: ValidationHandler = ( value ) => {
        if ( !value ) {
            return t( 'login.email_required' );
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

        formErrors.appendError( validateEmail( loginRequest.email as string ) );
        formErrors.appendError( validatePassword( loginRequest.password as string ) );

        handleFormError( formErrors );

        if ( formErrors.hasErrors() ) {
            return;
        }

        props.login( loginRequest );
    };

    return <Form onSubmit={(r) => onSubmit(r)} form={form} className={"relative lg:w-1/2"}>
        <Field.FormInput name="email"
               type="text"
               placeholder={t('login.email_placeholder')}
               className="col-span-12"
               validators={ [validateEmail] }
        />
        <Field.FormInput name="password"
               type="password"
               placeholder="Password"
               className="col-span-12"
               validators={ [validatePassword] }
        />
        <div className="flex justify-end w-full text-sm col-span-12">
            <a className="link link-neutral link-hover"> {t('login.forgot_password')} </a>
        </div>
        <Button
            type={"primary"}
            htmlType={"submit"}
            className="col-span-12"
        >
            {t('login.Submit')}
        </Button>
    </Form>;
};

export default LoginForm;