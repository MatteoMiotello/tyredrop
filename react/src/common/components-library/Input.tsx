import React, {ChangeEventHandler, PropsWithChildren, useState} from "react";
import {PropsWithValidators, ValidationHandler} from "../validation/validators";

interface FormControlProps extends PropsWithChildren {
    className?: string;
}

interface InputProps extends PropsWithValidators {
    type: string;
    name: string;
    placeholder: string;
    required?: boolean;
    labelText?: string | undefined | null;
    topRightLabelText?: string;
    bottomLeftLabelText?: string;
    bottomRightLabelText?: string;
    error?: string;
    size?: 'lg' | 'md' | 'sm' | 'xs' | undefined;
    value?: string;
}

interface FormInputProps extends FormControlProps, InputProps {

}

interface InputGroupProps extends PropsWithChildren {
    size?: 'lg' | 'md' | 'sm' | 'xs' | undefined;
}

const FormControl: React.FC<FormControlProps> = (props) => {
    return <div className={"form-control " + (props.className ?? '')}>
        {props.children}
    </div>;
};

const Input: React.FC<InputProps> = (props) => {
    const [error, setError] = useState<string | null>(props.error as string);

    let bottomLeft = null;
    let classes = null;
    const onChange: ChangeEventHandler<HTMLInputElement> = (event) => {
        const value = event.target.value;

        if (!props.validators) {
            return;
        }

        props.validators?.every((validator: ValidationHandler) => {
            const error = validator(value);

            setError(error);

            if (error) {
                return false;
            }

            return true;
        });
    };

    if (error) {
        bottomLeft = <span className="label-text-alt text-error">{error}</span>;
        classes = ' input-error ';
    }

    const getSizeClassName = (size: string | undefined): string => {
        switch (size) {
            case 'lg':
                return 'input-lg';
            case 'md':
                return 'input-md';
            case 'sm':
                return 'input-sm';
            case 'xs':
                return 'input-xs';
            default:
                return '';
        }
    };

    classes += ' ' + getSizeClassName(props.size);

    return <>
        {
            props.labelText || props.topRightLabelText ?
                <label className="label">
                    {props.labelText && <span className="label-text">{props.labelText}</span>}
                    {props.topRightLabelText && <span className="label-text-alt">{props.topRightLabelText}</span>}
                </label> :
                ''
        }
        <input type={props.type}
               name={props.name}
               placeholder={props.placeholder}
               className={"w-full input input-bordered " + classes}
               required={props.required}
               onChange={onChange}
               onFocus={onChange}
               value={props.value}
        />
        {
            bottomLeft || props.bottomRightLabelText ?
                <label className="label">
                    {props.bottomLeftLabelText && <span className="label-text-alt">{props.bottomLeftLabelText}</span>}
                    {bottomLeft}
                </label>
                :
                ''
        }
    </>;
};

const FormInput: React.FC<FormInputProps> = (props) => {
    return <FormControl className={props.className}>
        <Input{...props}/>
    </FormControl>;
};

const InputGroup: React.FC<InputGroupProps> = (props) => {
    const getSizeClass = (size: string | undefined): string => {
        switch (size) {
            case 'lg':
                return 'input-group-lg';
            case 'md':
                return 'input-group-md';
            case 'sm':
                return 'input-group-sm';
            case 'xs':
                return 'input-group-xs';
            default:
                return '';
        }
    };

    return <div className={"input-group " + getSizeClass(props.size)}>
        {props.children}
    </div>;
};

const Field = {
    FormControl,
    Input,
    FormInput,
    InputGroup
};

export default Field;