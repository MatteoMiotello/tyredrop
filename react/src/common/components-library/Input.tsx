import React, {ChangeEventHandler, useState} from "react";

export type ValidationHandler = (value: string | null ) => string | null

interface InputProps {
    type: string
    name: string
    placeholder: string
    required?: boolean
    labelText?: string
    topRightLabelText?: string
    bottomLeftLabelText?: string
    bottomRightLabelText?: string
    className?: string
    validate?: ValidationHandler
}


const Input: React.FC<InputProps> = (props) => {
    const [error, setError] = useState<string | null>(null);

    let bottomLeft = null;
    let classes = null;
    const onChange: ChangeEventHandler<HTMLInputElement> = (event) => {
        const value = event.target.value;

        if ( !props.validate) {
            return;
        }

        setError( props.validate( value ) );
    };

    if ( error ) {
        bottomLeft = <span className="label-text-alt text-error">{error}</span>;
        classes = ' input-error ';
    }

    return <div
        className={"my-2 form-control mx-2 " + (props.className ?? '')}>
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
               className={"input input-bordered" + classes}
               required={props.required}
               onChange={onChange}
               onFocus={onChange}
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
    </div>;
};

export default Input;