import React, {useEffect, useState} from "react";
import {useSelector} from "react-redux";
import {z} from "zod";
import {RegisterRequest} from "../../../common/backend/requests/register-request";
import Form, {FormErrors, FormSubmitHandler, useForm} from "../../../common/components-library/Form";
import Field from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import {useTranslation} from "react-i18next";
import {ValidationHandler} from "../../../common/validation/validators";
import {zodParser} from "../../../common/validation/zod-parser";
import {selectAuthStatus} from "../store/auth-selector";

interface RegisterInput {
    email: string;
    password: string;
    repeat_password: string;
    name: string | null;
    surname: string | null;
}

type RegisterFormProps = {
    onSuccess?: () => void
    register: (request: RegisterRequest) => void
}

const RegisterForm: React.FC<RegisterFormProps> = (props: RegisterFormProps) => {
    const [form, handleError] = useForm();
    const {t} = useTranslation();
    const [currentPassword, setCurrentPassword] = useState<string | null>(null);
    const authStatus = useSelector(selectAuthStatus);

    useEffect(() => {
        if (authStatus.status == 'error') {
            let error = authStatus.error;
            if (authStatus.error && (typeof authStatus.error == 'number' && authStatus.error >= 5000)) {
                error = t('register.email_already_used');
            }

            handleError(error);
        }

        if (authStatus.status == 'fullfilled' && props.onSuccess) {
            props.onSuccess();
        }
    }, [authStatus]);

    const validateEmail: ValidationHandler = (value: string | null): string | null => {
        const email = z.string()
            .nonempty({message: t('login.email_required') as string})
            .email({message: t('register.invalid_email') as string});

        return zodParser(email, value);
    };

    const validatePassword: ValidationHandler = (value: string | null): string | null => {
        setCurrentPassword(value);

        const passwordSchema = z.string()
                .nonempty({message: t('login.password_required') as string})
                .min(8, {message: t('register.password_requirement') as string});


        return zodParser(passwordSchema, value);

    };

    const validateRepeatPassword: ValidationHandler = (value: string | null): string | null => {
        if (!value) {
            return t('register.repeat_password_required');
        }

        if (value != currentPassword) {
            return t('register.password_are_not_equal');
        }


        return null;
    };


    const onSubmit: FormSubmitHandler<RegisterInput> = (registerRequest: RegisterInput) => {
        const formErrors = new FormErrors();

        formErrors.appendError(validateEmail(registerRequest.email));
        formErrors.appendError(validatePassword(registerRequest.password));
        formErrors.appendError(validateRepeatPassword(registerRequest.repeat_password));

        if (formErrors.hasErrors()) {
            handleError(formErrors);
            return;
        }

        props.register({
            email: registerRequest.email,
            password: registerRequest.password,
            name: registerRequest.name,
            surname: registerRequest.surname,
        });
    };

    return <Form form={form} onSubmit={onSubmit} className="lg:w-1/2 relative">
        <Field.FormInput name="email"
               type="text"
               placeholder={t('login.email_placeholder')}
               className="col-span-12"
               validators={[validateEmail]}
        />
        <Field.FormInput name="password"
               type="password"
               placeholder={t('login.password_placeholder')}
               className="col-span-6"
               validators={[validatePassword]}
        />
        <Field.FormInput name="repeat_password"
               type="password"
               placeholder={t('register.repeat_password_placeholder')}
               className="col-span-6"
               validators={[validateRepeatPassword]}
        />
        <Field.FormInput name="name"
               type="text"
               placeholder={t('register.name_placeholder')}
               className="col-span-6"
        />
        <Field.FormInput name="surname"
               type="text"
               placeholder={t('register.surname_placeholder')}
               className="col-span-6"
        />
        <div className="flex justify-end w-full text-sm col-span-12">
            <a className="link link-neutral link-hover"> {t('login.forgot_password')} </a>
        </div>
        <Button
            type={"primary"}
            htmlType={"submit"}
            className="col-span-12"
        >
            {t('register.submit_button')}
        </Button>
    </Form>;
};

export default RegisterForm;