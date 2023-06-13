import React, {PropsWithChildren} from "react";

type ButtonType = 'primary' | 'secondary' | 'accent' | 'ghost' | 'link' | 'info' | 'success' | 'error' | undefined;

interface ButtonProps extends PropsWithChildren {
    type?: ButtonType;
    size?: 'lg' | 'sm' | 'xs' | undefined;
    htmlType?: 'submit' | 'reset' | 'button' | undefined;
    outline?: boolean;
    className?: string;
    onClick?: () => void;
}

const getClassType = (type: ButtonType) => {
    switch (type) {
        case 'primary':
            return 'btn-primary';
        case 'secondary':
            return 'btn-secondary';
        case 'accent':
            return 'btn-accent';
        case 'ghost':
            return 'btn-ghost';
        case 'link':
            return 'btn-link';
        case 'info':
            return 'btn-info';
        case 'success':
            return 'btn-success';
        case 'error':
            return 'btn-error';
        default:
            return '';
    }
};

const getSizeClass = (size: string): string => {
    switch (size) {
        case 'lg':
            return 'btn-lg';
        case 'sm':
            return 'btn-sm';
        case 'xs':
            return 'btn-xs';
        default:
            return '';
    }
};

const Button: React.FC<ButtonProps> = (props) => {
    let classes = 'btn ';

    if (props.className) {
        classes += ' ' + props.className;
    }

    if (props.type) {
        classes += ' ' + getClassType(props.type);
    }

    if (props.outline) {
        classes += ' btn-outline';
    }

    if (props.size) {
        classes += ' ' + getSizeClass(props.size);
    }

    return <button className={classes} type={props.htmlType} onClick={props?.onClick}>
        {props.children}
    </button>;
};

export default Button;