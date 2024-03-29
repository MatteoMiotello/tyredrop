import React, {ChangeEventHandler, InputHTMLAttributes, PropsWithChildren, useRef, useState} from "react";
import {twMerge} from "tailwind-merge";
import {PropsWithValidators, ValidationHandler} from "../validation/validators";

interface FormControlProps extends PropsWithChildren {
    className?: string;
}

type InputProps = {
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
    defaultValue?: string;
    value?: any;
    onValueChange?: (value: any ) => void
} & PropsWithValidators & InputHTMLAttributes<HTMLInputElement>

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
    const ref = useRef( props.defaultValue );
    const [error, setError] = useState<string | null>(props.error as string);

    let bottomLeft = null;
    let classes = null;
    const onChange: ChangeEventHandler<HTMLInputElement> = (event) => {
        const value = event.target.value;
        if ( props.onValueChange ) {
            props.onValueChange(value);
        }

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

    classes += ' ' + props.className;

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
               value={props.value}
               className={"w-full input input-bordered " + classes}
               required={props.required}
               onChange={onChange}
               onFocus={onChange}
               defaultValue={props.defaultValue}
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

type LabelProps = {
    className?: string
} & PropsWithChildren

const Label: React.FC<LabelProps> = ({children, className}) => {
    const classNames = twMerge(
        'label-text',
        className
    );

    return <label className="label">
        <div className={classNames}>
            {children}
        </div>
    </label>;
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
    InputGroup,
    Label
};

export default Field;