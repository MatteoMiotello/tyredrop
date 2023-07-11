import React from "react";
import {useLoaderData} from "react-router-dom";
import {FetchOrderQuery, Order} from "../../__generated__/graphql";
import Panel from "../../common/components-library/Panel";
import OrderRowsTable from "./components/OrderRowsTable";
import OrderStatusBadge from "./components/OrderStatusBadge";
import {calculateTotal} from "./utils";

const OrderDetailsPage: React.FC = () => {
    const order = useLoaderData() as FetchOrderQuery;

    return <main className="p-4 grid grid-flow-row md:grid-cols-12 gap-4">
        <h1 className="col-span-12 text-center text-3xl my-10">
            Ordine n. #{order.order.id}
        </h1>
        <Panel className="col-span-8 row-span-2">
            <h3 className="font-semibold">Prodotti acquistati</h3>
            <OrderRowsTable order={order}></OrderRowsTable>
        </Panel>
        <Panel className="col-span-4 flex flex-col">
            <h3 className="font-semibold">Totale Ordine</h3>
            <div className="w-full my-auto text-center text-6xl font-bold text-primary">{calculateTotal(order.order as Order)}</div>
        </Panel>
        <Panel className="col-span-4 flex flex-col">
            <h3 className="font-semibold">Stato dell'ordine</h3>
            <span className="font-bold text-secondary mx-auto text-4xl"><OrderStatusBadge className="badge-lg" status={order.order.status}/></span>
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
                <li><strong>Indirizzo:</strong> {order.order.addressLine1}</li>
                {order.order.addressLine2 && <li><strong>Indirizzo 2:</strong> {order.order.addressLine2}</li>}
                <li><strong>Città:</strong> {order.order.city}</li>
                <li><strong>Paese:</strong> {order.order.country}</li>
                <li><strong>Provincia:</strong> {order.order.province}</li>
                <li><strong>CAP:</strong> {order.order.postalCode}</li>
            </ul>
        </Panel>
    </main>;
};

export default OrderDetailsPage;