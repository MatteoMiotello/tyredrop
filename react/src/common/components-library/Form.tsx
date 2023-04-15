import React, {FormEventHandler, PropsWithChildren, useState} from "react";
import Alert from "./Alert";
import {HookHandler} from "vite";


type FormProps = {
    className?: string
    onSubmit: FormEventHandler
    form: FormProperties
} & PropsWithChildren

type FormProperties = {
    error?: string | null
}

export const useForm: HookHandler<any> = () => {
    const [form, setForm] = useState<FormProperties>({error: null});

    const handleFormError = (error: string) => {
        setForm({error: error});
    };

    return [form, handleFormError];
};

const Form: React.FC<FormProps> = (props: FormProps) => {
    return <form
            className={"w-full inline-grid grid-cols-12 " + (props.className ?? '')}
            onSubmit={props.onSubmit}
        >
            {props.form.error ?
                <Alert
                    type="error"
                    className="col-span-12 lg:col-start-4 lg:col-span-6 w-auto"
                >
                    {props.form.error}
                </Alert> : ''
            }
            {props.children}
        </form>;
};

export default Form;