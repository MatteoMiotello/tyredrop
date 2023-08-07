import React from "react";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faDotCircle} from "@fortawesome/free-solid-svg-icons";

type AvailabilityBadgeProps = {
    quantity: number
}
const AvailabilityBadge: React.FC<AvailabilityBadgeProps> = ({quantity}) => {
    let config = {
        textColor: 'text-success',
        message: 'Disponibile (20+)'
    };

    if ( quantity > 10 && quantity < 20 ) {
        config = {
            textColor: 'text-info',
            message: `Quantità limitata (${quantity})`
        };
    } else if ( quantity < 10 && quantity > 4 ) {
        config = {
            textColor: 'text-warning',
            message: `Quantità ridotta (${quantity})`
        };
    } else if ( quantity < 4 ) {
        config = {
            textColor: 'text-error',
            message: `Non disponibile (${quantity})`
        };
    }

    return <span className={'text-sm inline-block ' + config.textColor}> <FontAwesomeIcon icon={faDotCircle}/> {config.message} </span>;
};

export default AvailabilityBadge;