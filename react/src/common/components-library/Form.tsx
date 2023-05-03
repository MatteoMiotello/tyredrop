import React, {Children, PropsWithChildren, ReactElement, useEffect, useState} from "react";
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

    return [form, handleFormError];
};

const Form: React.FC<FormProps> = (props: FormProps) => {
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

        const formData = new FormData(e.currentTarget);
        const values: any = {};

        formData.forEach((value, key) => {
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            values[key] = value;
        });

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
                    const value = values[childProps.name];

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

        props.onSubmit(values);

        if (props.onSuccess) {
            props.onSuccess();
        }
    };

    return <form
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
};

export default Form;