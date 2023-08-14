import React, {useEffect} from "react";
import Moment from "react-moment";
import {useLoaderData, useNavigate} from "react-router-dom";
import {FetchOrderQuery, Order, OrderStatus} from "../../__generated__/graphql";
import Panel from "../../common/components-library/Panel";
import {Badge, Button, useModal} from "../../common/components/shelly-ui";
import {Currency} from "../../common/utilities/currency";
import OrderRowsTable from "./components/OrderRowsTable";
import OrderStatusBadge from "./components/OrderStatusBadge";
import OrderSupportModal from "./components/OrderSupportModal";

const OrderDetailsPage: React.FC = () => {
    const order = useLoaderData() as FetchOrderQuery;
    const modal = useModal();
    const navigate = useNavigate();

    useEffect(() => {
        if (order.order.status == OrderStatus.NotCompleted) {
            navigate(`/order/checkout/${order.order.id}`);
        }
    }, [order]);

    return <main className="p-4 grid grid-flow-row md:grid-cols-12 gap-4">
        <OrderSupportModal modal={modal} order={order.order as Order}/>
        <div className="col-start-11 col-span-2 flex justify-end">
            <Button size="sm" buttonType="warning" onClick={modal.open}>
                Richiedi assistenza
            </Button>
        </div>
        <div className="col-span-12 text-center my-10">
            <h1 className=" text-3xl ">
                Ordine n. #{order.order.orderNumber}
            </h1>
            <Moment className="text-neutral">{order.order.createdAt}</Moment>
        </div>
        <Panel className="col-span-8 row-span-2">
            <h3 className="font-semibold">Prodotti acquistati</h3>
            <OrderRowsTable order={order}></OrderRowsTable>
        </Panel>
        <Panel className="col-span-4 flex flex-col">
            <h3 className="font-semibold">Totale Ordine</h3>
            <div
                className="w-full my-auto text-center text-6xl font-bold text-primary">{Currency.defaultFormat(
                order.order.priceAmountTotal,
                order.order.currency.iso_code
            )}</div>
            <div className="text-sm mt-4">
                <div>
                    Totale IVA (22%): {Currency.defaultFormat(order.order.taxesAmount, order.order.currency.iso_code)}
                </div>
                <div>
                    Totale senza IVA: {Currency.defaultFormat(order.order.priceAmount, order.order.currency.iso_code)}
                </div>
            </div>
        </Panel>
        <Panel className="col-span-4 flex flex-col">
            <h3 className="font-semibold">Stato dell'ordine</h3>
            <span className="font-bold text-secondary mx-auto text-4xl"><OrderStatusBadge className="badge-lg"
                                                                                          status={order.order.status}/></span>
        </Panel>
        <Panel className="col-span-6">
            <h3 className="font-semibold">Dati di fatturazione</h3>
            <ul className="ml-2">
                <li><strong>Nome:</strong> {order.order.userBilling.name}</li>
                <li><strong>Cognome:</strong> {order.order.userBilling.surname}</li>
                <li><strong>Codice fiscale:</strong> {order.order.userBilling.fiscalCode}</li>
                <li><strong>Partita IVA:</strong> {order.order.userBilling.vatNumber}</li>
                <li><strong>Tipo entità legale:</strong> {order.order.userBilling.legalEntityType.name}</li>
            </ul>
        </Panel>
        <Panel className="col-span-6">
            <h3 className="font-semibold">Indirizzo di spedizione</h3>
            <ul className="ml-2">
                <li><strong>Nome:</strong> {order.order.addressName}</li>
                <li><strong>Indirizzo:</strong> {order.order.addressLine1}</li>
                {order.order.addressLine2 && <li><strong>Indirizzo 2:</strong> {order.order.addressLine2}</li>}
                <li><strong>Città:</strong> {order.order.city}</li>
                <li><strong>Paese:</strong> {order.order.country}</li>
                <li><strong>Provincia:</strong> {order.order.province}</li>
                <li><strong>CAP:</strong> {order.order.postalCode}</li>
            </ul>
        </Panel>
        {
            order.order?.payment &&
            <Panel className="col-span-6">
                <h3 className="font-semibold">Dati del pagamento</h3>
                <p className="font-medium my-4"> Totale: { Currency.defaultFormat( order.order.payment.amount, order.order.currency.iso_code ) } </p>
                <ul className="">
                    <li > Metodo: <span className="font-medium"> {order.order.payment.userPaymentMethod.paymentMethod.name}</span> </li>
                    {
                        order.order.payment.userPaymentMethod.paymentMethod.receiver &&
                        <li> Beneficiario: {order.order.payment.userPaymentMethod.paymentMethod.receiver} </li>
                    }
                    {
                        order.order.payment.userPaymentMethod.paymentMethod.bank_name &&
                        <li> Istituto: {order.order.payment.userPaymentMethod.paymentMethod.bank_name} </li>
                    }
                    {
                        order.order.payment.userPaymentMethod.paymentMethod.bank_name?.length &&
                        <li> IBAN: {order.order.payment.userPaymentMethod.paymentMethod.iban} </li>
                    }
                    {
                        order.order.payment.userPaymentMethod.paymentMethod.receiver &&
                        <li> Causale: <Badge>titw_order_#{order.order.orderNumber}</Badge> </li>
                    }
                </ul>
            </Panel>
        }
    </main>
        ;
};

export default OrderDetailsPage;