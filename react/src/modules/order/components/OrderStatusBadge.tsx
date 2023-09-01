import React from "react";
import {OrderStatus} from "../../../__generated__/graphql";
import {Badge} from "../../../common/components/shelly-ui";

type OrderStatusBadgeProps = {
    status: OrderStatus
    className?: string
}

export const badgeConfig: { [key in OrderStatus]: { title: string, color: string } } = {
    NOT_COMPLETED: {
        title: "Da completare",
        color: "#A0A0A0"
    },
    CANCELED: {
        title: "Cancellato",
        color: "#676767"
    },
    CONFIRMED: {
        title: "Confermato",
        color: "#66FF66"
    },
    DELIVERED: {
        title: "Spedito",
        color: "#6666FF"
    },
    TO_PAY: {
        title: "In attesa di pagamento",
        color: "#f33939",
    },
    NEW: {
        title: "Nuovo",
        color: "#FFC266"
    },
    REJECTED: {
        title: "Rifiutato",
        color: "#cc0000"
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