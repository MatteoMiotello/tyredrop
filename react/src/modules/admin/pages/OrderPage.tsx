import React from "react";
import {useQuery} from "../../../common/backend/graph/hooks";
import {FETCH_ALL_ORDERS} from "../../../common/backend/graph/query/order";
import {AllOrdersQuery, AllOrdersQueryVariables} from "../../../__generated__/graphql";
import Panel from "../../../common/components-library/Panel";
import OrderAdminTable from "../components/Order/OrderAdminTable";

const OrderPage: React.FC =  () => {
    const query = useQuery<AllOrdersQuery, AllOrdersQueryVariables>( FETCH_ALL_ORDERS, {
        variables: {
            pagination: {
                offset: 0,
                limit: 20
            },
            ordering: [
                {
                    column: 'created_at',
                    desc: true
                }
            ]
        }
    } );

    return <main>
        <Panel>
            <Panel.Title>
                Tutti gli ordini
            </Panel.Title>
            {
                query.data?.allOrders && <OrderAdminTable query={query}/>
            }
        </Panel>
    </main>;
};

export default OrderPage;