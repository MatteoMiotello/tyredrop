import React, {useState} from "react";
import Form, {FormErrors, FormSubmitHandler, useForm} from "../../../common/components-library/Form";
import Input, {ValidationHandler} from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import {useTranslation} from "react-i18next";


interface RegisterRequest {
    username: string
    password: string
    repeatPassword: string
    name: string | null
    surname: string | null
}
const RegisterForm: React.FC = () => {
    const [ form, handleError ] = useForm();
    const {t} = useTranslation();
    const [ currentPassword, setCurrentPassword ] = useState<string|null>( null );
    const validateEmail: ValidationHandler = ( value: string | null ): string | null => {
        if ( !value ) {
            return t( 'login.email_required' );
        }

        return null;
    };

    const validatePassword: ValidationHandler = ( value: string | null ): string | null => {
        setCurrentPassword( value );
        if ( !value ) {
            return t( 'login.password_required' );
        }

        return null;
    };

    const validateRepeatPassword: ValidationHandler = ( value: string | null ): string | null => {
        if ( !value ) {
            return  t( 'register.repeat_password_required' );
        }

        if ( value != currentPassword ) {
            return t( 'register.password_are_not_equal' );
        }


        return null;
    };


    const onSubmit: FormSubmitHandler<RegisterRequest> = ( registerRequest: RegisterRequest ) => {
        const formErrors = new FormErrors();

        formErrors.appendError( validateEmail( registerRequest.username ) );
        formErrors.appendError( validatePassword( registerRequest.password ) );
        formErrors.appendError( validateRepeatPassword( registerRequest.repeatPassword ) );

        if ( formErrors.hasErrors() ) {
            handleError( formErrors );
        }
    };

    return <Form form={form} onSubmit={onSubmit} className="lg:w-1/2">
        <Input name="email"
               type="text"
               placeholder={t('login.email_placeholder')}
               className="col-span-12"
               validators={ [validateEmail] }
        />
        <Input name="password"
               type="password"
               placeholder="Password"
               className="col-span-6"
               validators={ [validatePassword] }
        />
        <Input name="repeat_password"
               type="password"
               placeholder={ t( 'register.repeat_password_placeholder' ) }
               className="col-span-6"
               validators={ [validateRepeatPassword] }
        />
        <Input name="name"
               type="text"
               placeholder={ t( 'register.name_placeholder' ) }
               className="col-span-6"
        />
        <Input name="surname"
               type="text"
               placeholder={ t( 'register.surname_placeholder' ) }
               className="col-span-6"
        />
        <div className="flex justify-between w-full text-sm col-span-12">
            <a className="link link-neutral link-hover"
               href={"/auth/login"}> {t('register.login_label')} </a>
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

export default RegisterForm;