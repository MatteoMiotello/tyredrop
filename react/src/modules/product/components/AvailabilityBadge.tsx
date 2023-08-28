import React from "react";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faDotCircle} from "@fortawesome/free-solid-svg-icons";

type AvailabilityBadgeProps = {
    quantity: number
}
const AvailabilityBadge: React.FC<AvailabilityBadgeProps> = ({quantity}) => {
    let config = {
        textColor: '#2ae800',
        message: 'Disponibile'
    };

    if ( quantity > 10 && quantity < 20 ) {
        config = {
            textColor: '#4287f5',
            message: `Quantità limitata `
        };
    } else if ( quantity < 10 && quantity > 4 ) {
        config = {
            textColor: '#ff2424',
            message: `Quantità ridotta `
        };
    } else if ( quantity < 4 ) {
        config = {
            textColor: '#ff2424',
            message: `Non disponibile `
        };
    }

    return <span className={'text-sm font-medium inline-block ' } style={{color: config.textColor}}> <FontAwesomeIcon icon={faDotCircle}/> {config.message} <br/> {quantity > 20 ? '(+20)' : `(${quantity})`} </span>;
};

export default AvailabilityBadge;
