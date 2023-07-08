import React from "react";
import {useQuery} from "@apollo/client";
import {Order} from "../../__generated__/graphql";
import {FETCH_USER_ORDERS} from "../../common/backend/graph/query/order";
import {useAuth} from "../auth/hooks/useAuth";
import Spinner from "../../common/components/Spinner";
import {useToast} from "../../hooks/useToast";
import OrderTable from "../order/components/OrderTable";

const UserOrdersPage: React.FC = () => {
    const auth = useAuth();
    const {setError} = useToast();
    const {data, loading, error} = useQuery(FETCH_USER_ORDERS, {
        variables: {
            userId: auth.user?.user?.userID as string,
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

    return <main className="p-4">
        <OrderTable orders={data?.userOrders as Order[]}></OrderTable>
    </main>;
};

export default UserOrdersPage;