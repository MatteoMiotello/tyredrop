import React, {Children, PropsWithChildren, ReactElement, forwardRef, useEffect, useState} from "react";
import {ValidationHandler} from "../validation/validators";
import Alert from "./Alert";
import {HookHandler} from "vite";
import {faTriangleExclamation} from "@fortawesome/free-solid-svg-icons";

export class FormErrors {
    public errors: string[] = [];

    appendError(error: string | null) {
        if (!error) {
            return;
        }

        this.errors.push(error);
    }

    hasErrors(): boolean {
        return this.errors.length > 0;
    }

    resetErrors(): void {
        this.errors = [];
    }
}

export type FormOnSuccessHandler = () => void
export type FormSubmitHandler<T> = (values: T) => void

interface FormProps<T = any> extends PropsWithChildren {
    className?: string;
    onSubmit: FormSubmitHandler<T>;
    onSuccess?: FormOnSuccessHandler;
    form: FormProperties;
}

type FormProperties = {
    formError: FormErrors
}

export const useForm: HookHandler<any> = () => {
    const [form, setForm] = useState<FormProperties>({formError: new FormErrors()});

    const handleFormError = (formError: FormErrors | string | null) => {
        if (!formError) {
            return;
        }

        if (!(formError instanceof FormErrors)) {
            const sFormError = formError as string;

            formError = new FormErrors();
            formError.appendError(sFormError);
        }


        setForm({formError: formError});
    };

    return {form, handleFormError};
};

const Form = forwardRef<HTMLFormElement, FormProps>( (props, ref) => {
    const [form, setForm] = useState(props.form);

    useEffect( () => {
        if ( form.formError.hasErrors() ) {
            return;
        }

        setForm(props.form);
    }, [props.form] );

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        form.formError.resetErrors();
        e.preventDefault();

        const formData = new FormData( e.currentTarget );
        const data: { [key: string]: any } = {};

        for (const pair of formData.entries()) {
            const key = pair[0];
            const value = pair[1];

            if (key.includes('[') && key.includes(']')) {
                const fieldName = key.substring(0, key.indexOf('['));
                const matches = key.match(/\[(.*?)\]/);
                const index = matches ? matches[1] : 0;

                if (!Object.prototype.hasOwnProperty.call(data, fieldName)) {
                    data[fieldName] = [];
                }

                data[fieldName][index] = value;
            } else if (key.includes('{') && key.includes('}')) {
                const objFieldName = key.substring(0, key.indexOf('{'));
                const matches = key.match(/\[(.*?)\]/);
                const objKey = matches ? matches[1] : 0;

                if (!Object.prototype.hasOwnProperty.call(data, objFieldName)) {
                    data[objFieldName] = {};
                }

                data[objFieldName][objKey] = value;
            } else {
                data[key] = value;
            }
        }

        if (props.children) {
            const arrayChildren = Children.toArray(props.children);

            const formErrors = new FormErrors();
            arrayChildren.forEach((child) => {

                child = child as ReactElement;

                const childProps = child.props;
                if (!childProps) {
                    return;
                }

                if (!childProps.validators) {
                    return;
                }

                if (!childProps.name) {
                    return;
                }

                childProps.validators?.forEach((validator: ValidationHandler) => {
                    const value = data[childProps.name];

                    if ( value === undefined ) {
                        return;
                    }

                    const error = validator(value);

                    if (error) {
                        formErrors.appendError(error);
                    }
                });

            });

            if (formErrors.hasErrors()) {
                setForm({formError: formErrors});
            }
        }

        if (form.formError.hasErrors()) {
            return;
        }

        props.onSubmit(data);

        if (props.onSuccess) {
            props.onSuccess();
        }
    };

    return <form
        ref={ref}
        className={"w-full inline-grid grid-cols-12 gap-4 " + (props.className ?? '')}
        onSubmit={(e) => onSubmit(e)}
    >
        {form.formError.errors.length > 0 ?
            <Alert
                type="error"
                className="col-span-12 w-auto text-sm"
                icon={faTriangleExclamation}
            >
                <h3 className="m-1"> Errore nel salvataggio </h3>
                {form.formError.errors.map((error, index) => {
                    return <li className="m-1" key={index}> {error} </li>;
                })}
            </Alert> : ''
        }
        {props.children}
    </form>;
} );

Form.displayName = 'Form';

export default Form;