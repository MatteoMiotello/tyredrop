import React from "react";
import {useQuery} from "@apollo/client";
import {useParams} from "react-router-dom";
import {Order} from "../../__generated__/graphql";
import {FETCH_USER_ORDERS} from "../../common/backend/graph/query/order";
import Panel from "../../common/components-library/Panel";
import Spinner from "../../common/components/Spinner";
import {useToast} from "../../hooks/useToast";
import OrderTable from "../order/components/OrderTable";

const UserOrdersPage: React.FC = () => {
    const params = useParams();
    const {setError} = useToast();
    const {data, loading, error} = useQuery(FETCH_USER_ORDERS, {
        variables: {
            userId: params.id as string,
        },
        fetchPolicy: "no-cache"
    });

    if (loading) {
        return <Spinner></Spinner>;
    }

    if (error) {
        setError('C\'è stato un errore nel caricamento');
        return null;
    }

    return <main className="">
        <Panel className="min-h-full">
            <h3 className="text-xl"> Tutti gli ordini </h3>
            {data?.userOrders == null && <span> Non e` stato effettuato alcun ordine </span>}
            {data?.userOrders && <OrderTable orders={data?.userOrders as Order[]}></OrderTable>}
        </Panel>
    </main>;
};

export default UserOrdersPage;