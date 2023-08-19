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

export function OrderBillingPanel(props: { order: FetchOrderQuery, className: string }) {
    return <Panel className={props.className}>
        <h3 className="font-semibold">Dati di fatturazione</h3>
        <ul className="ml-2">
            <li><strong>Nome:</strong> {props.order.order.userBilling.name}</li>
            <li><strong>Cognome:</strong> {props.order.order.userBilling.surname}</li>
            <li><strong>Codice fiscale:</strong> {props.order.order.userBilling.fiscalCode}</li>
            <li><strong>Partita IVA:</strong> {props.order.order.userBilling.vatNumber}</li>
            <li><strong>Tipo entità legale:</strong> {props.order.order.userBilling.legalEntityType.name}</li>
        </ul>
    </Panel>;
}

export function OrderAddressPanel(props: { order: FetchOrderQuery, className: string }) {
    return <Panel className={props.className}>
        <h3 className="font-semibold">Indirizzo di spedizione</h3>
        <ul className="ml-2">
            <li><strong>Nome:</strong> {props.order.order.addressName}</li>
            <li><strong>Indirizzo:</strong> {props.order.order.addressLine1}</li>
            {props.order.order.addressLine2 && <li><strong>Indirizzo 2:</strong> {props.order.order.addressLine2}</li>}
            <li><strong>Città:</strong> {props.order.order.city}</li>
            <li><strong>Paese:</strong> {props.order.order.country}</li>
            <li><strong>Provincia:</strong> {props.order.order.province}</li>
            <li><strong>CAP:</strong> {props.order.order.postalCode}</li>
        </ul>
    </Panel>;
}

export function OrderTotalPanel(props: { order: FetchOrderQuery, className: string }) {
    return <Panel className={props.className + ' flex flex-col'} >
        <h3 className="font-semibold">Totale Ordine</h3>
        <div
            className="w-full my-auto text-center text-6xl font-bold text-primary">{Currency.defaultFormat(
            props.order.order.priceAmountTotal,
            props.order.order.currency.iso_code
        )}</div>
        <div className="text-sm mt-4">
            <div>
                Totale IVA
                (22%): {Currency.defaultFormat(props.order.order.taxesAmount, props.order.order.currency.iso_code)}
            </div>
            <div>
                Totale senza
                IVA: {Currency.defaultFormat(props.order.order.priceAmount, props.order.order.currency.iso_code)}
            </div>
        </div>
    </Panel>;
}

export function OrderPaymentDetails(props: { order: FetchOrderQuery }) {
    if (!props.order.order.payment) {
        return null;
    }

    return <>
        <p className="font-medium my-4"> Totale: {Currency.defaultFormat(props.order.order.payment.amount, props.order.order.currency.iso_code)} </p>
        <ul className="">
            <li> Metodo: <span
                className="font-medium"> {props.order.order.payment.userPaymentMethod.paymentMethod.name}</span></li>
            {
                props.order.order.payment.userPaymentMethod.paymentMethod.receiver &&
                <li> Beneficiario: {props.order.order.payment.userPaymentMethod.paymentMethod.receiver} </li>
            }
            {
                props.order.order.payment.userPaymentMethod.paymentMethod.bank_name &&
                <li> Istituto: {props.order.order.payment.userPaymentMethod.paymentMethod.bank_name} </li>
            }
            {
                props.order.order.payment.userPaymentMethod.paymentMethod.bank_name?.length &&
                <li> IBAN: {props.order.order.payment.userPaymentMethod.paymentMethod.iban} </li>
            }
            {
                props.order.order.payment.userPaymentMethod.paymentMethod.receiver &&
                <li> Causale: <Badge>titw_order_#{props.order.order.orderNumber}</Badge></li>
            }
        </ul>
    </>;
}

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
        <OrderTotalPanel order={order} className="col-span-4"/>
        <Panel className="col-span-4 flex flex-col">
            <h3 className="font-semibold">Stato dell'ordine</h3>
            <span className="font-bold text-secondary mx-auto text-4xl"><OrderStatusBadge className="badge-lg"
                                                                                          status={order.order.status}/></span>
        </Panel>
        <OrderBillingPanel order={order} className="col-span-6"/>
        <OrderAddressPanel order={order} className="col-span-6"/>
        {
            order.order?.payment &&
            <Panel className="col-span-6">
                <h3 className="font-semibold">Dati del pagamento</h3>
                <OrderPaymentDetails order={order}/>
            </Panel>
        }
    </main>;
};

export default OrderDetailsPage;