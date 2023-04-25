import React, {useState} from "react";
import Form, {FormSubmitHandler, useForm} from "../../../common/components-library/Form";
import Input, {ValidationHandler} from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import {useTranslation} from "react-i18next";


interface RegisterRequest {
    username: string
    password: string
    repeatPassword: string
}
const RegisterForm: React.FC = () => {
    const [ form, handleError ] = useForm();
    const {t} = useTranslation();
    const [ currentPassword, setCurrentPassword ] = useState<string|null>( null );
    const validateUsername: ValidationHandler = ( value: string | null ): string | null => {
        if ( !value ) {
            return t( 'login.username_required' );
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

    const onSubmit: FormSubmitHandler<RegisterRequest> = () => {
        return;
    };

    return <Form form={form} onSubmit={onSubmit}>
        <Input name="username"
               type="text"
               placeholder={t('login.username_placeholder')}
               className="col-span-12 lg:col-start-4 lg:col-span-6"
               validate={ validateUsername }
        />
        <Input name="password"
               type="password"
               placeholder="Password"
               className="col-span-6 lg:col-start-4 lg:col-span-6"
               validate={ validatePassword }
        />
        <Input name="repeat_password"
               type="password"
               placeholder={ t( 'register.repeat_password_placeholder' ) }
               className="col-span-6 lg:col-start-4 lg:col-span-6"
               validate={ validateRepeatPassword }
        />
        <div className="flex justify-between w-full text-sm col-span-12 lg:col-start-4 lg:col-span-6">
            <a className="link link-neutral link-hover"
               href={"/auth/register"}> {t('register.login_label')} </a>
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

export default RegisterForm;