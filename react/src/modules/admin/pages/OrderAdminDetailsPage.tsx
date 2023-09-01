import React from "react";
import {Link, useLoaderData, useNavigate} from "react-router-dom";
import {FetchOrderQuery, Order, OrderRowsQuery, OrderRowsQueryVariables} from "../../../__generated__/graphql";
import {useQuery} from "../../../common/backend/graph/hooks";
import {ORDER_ROWS} from "../../../common/backend/graph/query/order";
import Panel from "../../../common/components-library/Panel";
import {useForm} from "../../../common/components/shelly-ui";
import {OrderAddressPanel, OrderBillingPanel, OrderPaymentDetails, OrderTotalPanel} from "../../order/OrderDetailsPage";
import OrderPaymentForm from "../components/Order/OrderPaymentForm";
import OrderRowAdminTable from "../components/Order/OrderRowAdminTable";
import OrderStatusForm from "../components/Order/OrderStatusForm";

const OrderAdminDetailsPage: React.FC = () => {
    const query = useLoaderData() as FetchOrderQuery;
    const rowsQuery = useQuery<OrderRowsQuery, OrderRowsQueryVariables>( ORDER_ROWS, {
        variables: {
            orderId: query.order.id
        }
    } );
    const navigate = useNavigate();

    const statusForm = useForm({
        onSuccess: () => navigate('.', {replace: true})
    });
    const paymentForm = useForm({
        onSuccess: () => navigate('.', {replace: true})
    });

    return <main className="grid grid-cols-12 gap-2">
        <Panel className="col-span-12">
            <Panel.Title>
                Ordine n: #{query.order.orderNumber}
            </Panel.Title>
            Utente: <Link className="link-accent" to={`/admin/user/${query.order.userBilling.id}`}> {query.order.userBilling.user.email } </Link>
        </Panel>
        <OrderAddressPanel order={query} className="col-span-4"/>
        <OrderBillingPanel order={query} className="col-span-4"/>
        <OrderTotalPanel order={query} className="col-span-4"/>
        <Panel className="col-span-6">
            <Panel.Title>
                Pagamento
            </Panel.Title>
            {
                query.order.payment ?
                    <OrderPaymentDetails order={query}/> :
                    <OrderPaymentForm form={paymentForm} order={query.order as Order}/>
            }
        </Panel>
        <Panel className="col-span-6">
            <Panel.Title>
                Stato dell'ordine
            </Panel.Title>
            <OrderStatusForm form={statusForm} order={query.order as Order}/>
        </Panel>
        <Panel className="col-span-12">
            <Panel.Title>
                Articoli
            </Panel.Title>
            {
                rowsQuery.data &&
                <OrderRowAdminTable query={rowsQuery} order={query.order as Order}/>
            }
        </Panel>
    </main>;
};

export default OrderAdminDetailsPage;