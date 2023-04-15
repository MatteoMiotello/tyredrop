import React, {PropsWithChildren} from "react";
import {IconDefinition} from "@fortawesome/fontawesome-svg-core";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";

type AlertType = 'info' | 'success' | 'warning' | 'error' | undefined

interface AlertProps extends PropsWithChildren {
    icon?: IconDefinition
    className?: string
    type?: AlertType
}

const getTypeCssClass = ( type: AlertType ): string => {
    switch ( type ) {
        case 'info':
            return 'alert-info';
        case 'success':
            return 'alert-success';
        case 'warning':
            return 'alert-warning';
        case 'error':
            return 'alert-error';
        default:
            return '';
    }
};

const Alert: React.FC<AlertProps> = (props: AlertProps) => {
    let classType = " alert-info";

    if ( props.type ) {
        classType = getTypeCssClass( props.type );
    }

    return <div className={"my-2 alert " + classType + ' ' + ( props.className ?? '') }>
        <div>
            {props.icon && <FontAwesomeIcon size="xl" className={"mx-2"} icon={props.icon}/>}
            <span>
                {props.children}
            </span>
        </div>
    </div>;
};

export default Alert;