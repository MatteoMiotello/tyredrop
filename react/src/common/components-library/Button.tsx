import React, {PropsWithChildren} from "react";

type ButtonType = 'primary' | 'secondary' | 'accent' | 'ghost' | 'link' | undefined;

interface ButtonProps extends PropsWithChildren {
    type?: ButtonType
    size?: 'lg' | 'sm' | 'xs' | undefined
    htmlType?: 'submit' | 'reset' | 'button' | undefined;
    outline?: boolean
    className?: string
}

const getClassType = ( type: ButtonType ) => {
    switch ( type ) {
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
        default:
            return '';
    }
};

const Button: React.FC<ButtonProps> = (props) => {
    let classes = 'btn my-2';

    if (props.className) {
        classes += ' ' + props.className;
    }

    if (props.type) {
        classes += ' ' + getClassType( props.type );
    }

    if (props.outline) {
        classes += ' btn-outline';
    }

    if (props.size) {
        classes += ' btn-' + props.size;
    }

    return <button className={classes} type={props.htmlType}>
        {props.children}
    </button>;
};

export default Button;