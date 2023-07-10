import React, {PropsWithChildren} from "react";

type BadgeType = 'primary' | 'secondary' | 'accent' | 'ghost' | undefined
interface BadgeProps extends PropsWithChildren {
    outline: boolean
    type?: BadgeType
    className?: string
    color?: string
}

const Badge: React.FC<BadgeProps> = ( props ) => {
    const getTypeClass = ( type: BadgeType ) => {
        switch ( type ) {
            case 'primary': return 'badge-primary';
            case 'secondary': return 'badge-secondary';
            case 'accent': return 'badge-accent';
            case 'ghost': return 'badge-ghost';
            default: return '';
        }
    };

    return <div
        style={ {backgroundColor: ( !props.outline ? props.color : undefined ), borderColor: props.color } }
        className={`badge  ${getTypeClass(props.type)}${props.outline && ' badge-outline '} ${props.className}`}>
        {props.children}
    </div>;
};

export default Badge;