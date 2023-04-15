import React, {PropsWithChildren} from "react";
import {IconDefinition} from "@fortawesome/fontawesome-svg-core";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";

interface AlertProps extends PropsWithChildren {
    icon?: IconDefinition
    className?: string
    type?: 'info' | 'success' | 'warning' | 'error' | undefined
}

const Alert: React.FC<AlertProps> = (props: AlertProps) => {
    let classType = " alert-info";

    if ( props.type ) {
        classType = " alert-" + props.type;
    }

    return <div className={"my-2 alert " + classType + ' ' + ( props.className ?? '') }>
        <div>
            {props.icon && <FontAwesomeIcon icon={props.icon}/>}
            <span>
                {props.children}
            </span>
        </div>
    </div>;
};

export default Alert;