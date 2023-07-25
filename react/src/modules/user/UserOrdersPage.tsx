import React from "react";
import {useParams} from "react-router-dom";
import {FETCH_USER_ORDERS} from "../../common/backend/graph/query/order";
import Panel from "../../common/components-library/Panel";
import OrderTable from "../order/components/OrderTable";
import {useToast} from "../../store/toast";
import {useQuery} from "../../common/backend/graph/hooks";

const UserOrdersPage: React.FC = () => {
    const params = useParams();
    const toastr = useToast();
    const query = useQuery(FETCH_USER_ORDERS, {
        variables: {
            userId: params.id as string,
            pagination: {
                limit: 10,
                offset: 0
            }
        },
        fetchPolicy: "no-cache"
    });

    return <main className="">
        <Panel className="min-h-full">
            <h3 className="text-xl"> Tutti gli ordini </h3>
            {query.data?.userOrders == null && <span> Non e` stato effettuato alcun ordine </span>}
            {query.data?.userOrders && <OrderTable query={query}></OrderTable>}
        </Panel>
    </main>;
};

export default UserOrdersPage;