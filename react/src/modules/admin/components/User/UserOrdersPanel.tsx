import React from "react";
import Panel from "../../../../common/components-library/Panel";
import OrderTable from "../../../order/components/OrderTable";
import {useQuery} from "../../../../common/backend/graph/hooks";
import {FetchOrdersQuery, FetchOrdersQueryVariables, User} from "../../../../__generated__/graphql";
import {FETCH_USER_ORDERS} from "../../../../common/backend/graph/query/order";

type UserOrdersPanelProps = {
    user: User
}
const UserOrdersPanel: React.FC<UserOrdersPanelProps> = ({user}) => {
    const orderQuery = useQuery<FetchOrdersQuery, FetchOrdersQueryVariables>( FETCH_USER_ORDERS, {
        variables: {
            userId: user?.id as string,
            pagination: {
                limit: 10,
                offset: 0
            },
            ordering: [
                {
                    column: 'created_at',
                    desc: true
                }
            ]
        },
    } );

    return <Panel>
        <Panel.Title>
            Ordini
        </Panel.Title>
        {orderQuery.data?.userOrders && <OrderTable query={orderQuery}/>}
    </Panel>;
};

export default UserOrdersPanel;