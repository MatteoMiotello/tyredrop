import React, { PropsWithChildren, useState} from "react";
import Alert from "./Alert";
import {HookHandler} from "vite";
import {faTriangleExclamation} from "@fortawesome/free-solid-svg-icons";

export class FormErrors {
    public errors: string[] = [];

    appendError( error: string | null ) {
        if ( !error ) {
            return;
        }
        
        this.errors.push( error );
    }

    hasErrors(): boolean {
        return this.errors.length > 0;
    }
}

export type FormSubmitHandler<T> = ( values: T ) => void

type FormProps<T = any> = {
    className?: string
    onSubmit: FormSubmitHandler<T>
    form: FormProperties
} & PropsWithChildren

type FormProperties = {
    formError: FormErrors
}

export const useForm: HookHandler<any> = () => {
    const [form, setForm] = useState<FormProperties>({formError: new FormErrors()});

    const handleFormError = (formError: FormErrors) => {
        setForm({formError: formError});
    };

    return [form, handleFormError];
};

const Form: React.FC<FormProps> = (props: FormProps) => {
    const onSubmit = ( e: React.FormEvent<HTMLFormElement> ) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);

        const values = {};

        formData.forEach( ( value, key ) => {
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            values[key] = value;
        } );

        if ( props.form.formError.errors.length > 0 ) {
            return;
        }

        props.onSubmit( values );
    };

    return <form
            className={"w-full inline-grid grid-cols-12 " + (props.className ?? '')}
            onSubmit={onSubmit}
        >
            {props.form.formError.errors.length > 0 ?
                <Alert
                    type="error"
                    className="col-span-12 lg:col-start-4 lg:col-span-6 w-auto text-sm"
                    icon={faTriangleExclamation}
                >
                    <h3 className="m-1"> Errore nel salvataggio </h3>
                    {props.form.formError.errors.map( (error, index) =>  {
                        return <li className="m-1" key={index}> {error} </li>;
                    } )}
                </Alert> : ''
            }
            {props.children}
        </form>;
};

export default Form;