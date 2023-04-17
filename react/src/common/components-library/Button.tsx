import React, {PropsWithChildren} from "react";

interface ButtonProps extends PropsWithChildren {
    type?: 'primary' | 'secondary' | 'accent' | 'ghost' | 'link' | undefined
    size?: 'lg' | 'sm' | 'xs' | undefined
    htmlType?: 'submit' | 'reset' | 'button' | undefined;
    outline?: boolean
    className?: string
}

const Button: React.FC<ButtonProps> = (props) => {
    let classes = 'btn my-2';

    if (props.className) {
        classes += ' ' + props.className;
    }

    if (props.type) {
        classes += ' btn-' + props.type;
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