import React from "react";
import {OrderStatus} from "../../../__generated__/graphql";
import Badge from "../../../common/components-library/Badge";

type OrderStatusBadgeProps = {
    status: OrderStatus
    className?: string
}

export const badgeConfig: { [key in OrderStatus]: { title: string, color: string } } = {
    CANCELED: {
        title: "Cancellato",
        color: "#FF6666"
    },
    CONFIRMED: {
        title: "Confermato",
        color: "#66FF66"
    },
    DELIVERED: {
        title: "Spedito",
        color: "#6666FF"
    },
    NEW: {
        title: "Nuovo",
        color: "#FFC266"
    },
    REJECTED: {
        title: "Rifiutato",
        color: "#A0A0A0"
    },
    RETURNED: {
        title: "Rimborsato",
        color: "#FFFF66"
    }
};
const OrderStatusBadge: React.FC<OrderStatusBadgeProps> = ( {status, className} ) => {
    if ( !badgeConfig[status] ) {
        return <Badge>{status}</Badge>;
    }

    return <Badge className={className} color={badgeConfig[status].color}>
        {badgeConfig[status].title}
    </Badge>;
};

export default OrderStatusBadge;