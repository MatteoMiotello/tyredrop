import {useLazyQuery} from "@apollo/client";
import React, {useEffect, useState} from "react";
import {
    Order,
    OrderStatus, PossibleOrderStatusesQuery, PossibleOrderStatusesQueryVariables,
} from "../../../__generated__/graphql";
import {POSSIBLE_ORDER_STATUSES} from "../../../common/backend/graph/query/order";
import {Select, SelectOption} from "../../../common/components/shelly-ui";
import {InputProps} from "../../../common/components/shelly-ui/Form";
import OrderStatusBadge from "./OrderStatusBadge";

type OrderStatusSelectProps = {
    order?: Order
} & InputProps
const OrderStatusSelect: React.FC<OrderStatusSelectProps> = ({order, ...props}) => {
    const [options, setOptions] = useState<SelectOption[]>(Object.values(OrderStatus).map((val) => ({
        value: val,
        title: val
    })));
    const [fetch, query] = useLazyQuery<PossibleOrderStatusesQuery, PossibleOrderStatusesQueryVariables>(POSSIBLE_ORDER_STATUSES, {
        fetchPolicy: 'no-cache'
    });

    useEffect(() => {
        if (order) {
            fetch({
                variables: {
                    orderId: order.id
                }
            });
        }
    }, [order]);

    useEffect(() => {
        if (query.data) {
            setOptions(query.data.possibleOrderStatuses.map((status: OrderStatus) => ({value: status, title: status})));
        }
    }, [query]);

    return <Select
        options={options}
        displayFn={(opt) => <OrderStatusBadge status={opt.value}/>}
        {...props}
    />;
};

export default OrderStatusSelect;