import React from "react";


interface InputProps {
    type: string
    name: string
    placeholder: string
    required?: boolean
    labelText?: string
    topRightLabelText?: string
    bottomLeftLabelText?: string
    bottomRightLabelText?: string
}

const Input: React.FC<InputProps> = (props) => {
    return <div className="form-control w-full">
        <label className="label">
            {props.labelText ?? <span className="label-text">{props.labelText}</span>}
            {props.topRightLabelText ?? <span className="label-text-alt">{props.topRightLabelText}</span>}
        </label>
        <input type={props.type} name={props.name} placeholder={props.placeholder} className="input input-bordered" required={ props.required }/>
        <label className="label">
            {props.bottomLeftLabelText ?? <span className="label-text-alt">{props.bottomLeftLabelText}</span>}
            {props.bottomRightLabelText ?? <span className="label-text-alt">{props.bottomRightLabelText}</span>}
        </label>
    </div>
}

export default Input