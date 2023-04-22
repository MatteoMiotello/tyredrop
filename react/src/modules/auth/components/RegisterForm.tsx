import React from "react";
import Form, {useForm} from "../../../common/components-library/Form";
import Input, {ValidationHandler} from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import {useTranslation} from "react-i18next";

const RegisterForm: React.FC = () => {
    const [ form, handleError ] = useForm();
    const {t} = useTranslation();
    const validateUsername: ValidationHandler = ( value: string | null ): string | null => {
        if ( !value ) {
            return ;
        }

        return ;
    };

    const validatePassword: ValidationHandler = ( value: string | null ): string | null => {
        return;
    };

    const validateRepeatPassword: ValidationHandler = ( value: string | null ): string | null => {
        return; 
    };

    return <Form form={form}>
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
        <Input name="repeat_password"
               type="password"
               placeholder={ t( 'register.repeat_password_placeholder' ) }
               className="col-span-12 lg:col-start-4 lg:col-span-6"
               validate={ validateRepeatPassword }
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

export default RegisterForm;