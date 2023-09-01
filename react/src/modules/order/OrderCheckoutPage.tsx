import { faChevronUp} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Disclosure} from "@headlessui/react";
import React, {useEffect} from "react";
import {useLoaderData, useNavigate} from "react-router-dom";
import {
    ConfirmOrderMutation, ConfirmOrderMutationVariables,
    FetchOrderQuery,
    OrderStatus
} from "../../__generated__/graphql";
import Panel from "../../common/components-library/Panel";
import {Button} from "../../common/components/shelly-ui";
import {Currency} from "../../common/utilities/currency";
import OrderRowsTable from "./components/OrderRowsTable";
import {useMutation} from "../../common/backend/graph/hooks";
import {CONFIRM_ORDER} from "../../common/backend/graph/mutation/order";

const OrderCheckoutPage: React.FC = () => {
    const order = useLoaderData() as FetchOrderQuery;
    const navigate = useNavigate();
    const [mutate] = useMutation<ConfirmOrderMutation, ConfirmOrderMutationVariables>(CONFIRM_ORDER);

    useEffect(() => {
        if (order.order.status != OrderStatus.NotCompleted) {
            navigate(`/order/details/${order.order.id}`);
        }
    }, [order]);

    return <main className="flex flex-col items-start md:flex-row gap-2 mt-2 mx-1">
        <Panel className="col-span-2 flex-1 ">
            <Panel.Title>
                Riepilogo dell'ordine
            </Panel.Title>
            <div className="flex flex-col gap-2">
                <Disclosure as="div" className="collapse">
                    {({open}) => (
                        <>
                            <Disclosure.Button
                                className="w-full flex justify-between collapse-title bg-base-200 items-center rounded-box">
                                <span>Prodotti acquistati</span>
                                <FontAwesomeIcon icon={faChevronUp} transform={{rotate: (open ? 0 : 180)}}/>
                            </Disclosure.Button>
                            <Disclosure.Panel className="m-4">
                                <OrderRowsTable order={order}/>
                            </Disclosure.Panel>
                        </>
                    )}
                </Disclosure>
                <Disclosure as="div" className="collapse">
                    {({open}) => (
                        <>
                            <Disclosure.Button
                                className="w-full flex justify-between collapse-title bg-base-200 items-center rounded-box">
                                <span>Indirizzo di spedizione</span>
                                <FontAwesomeIcon icon={faChevronUp} transform={{rotate: (open ? 0 : 180)}}/>
                            </Disclosure.Button>
                            <Disclosure.Panel className="m-4">
                                <ul className="ml-2">
                                    <li><strong>Nome:</strong> {order.order.addressName}</li>
                                    <li><strong>Indirizzo:</strong> {order.order.addressLine1}</li>
                                    {order.order.addressLine2 &&
                                        <li><strong>Indirizzo 2:</strong> {order.order.addressLine2}</li>}
                                    <li><strong>Città:</strong> {order.order.city}</li>
                                    <li><strong>Paese:</strong> {order.order.country}</li>
                                    <li><strong>Provincia:</strong> {order.order.province}</li>
                                    <li><strong>CAP:</strong> {order.order.postalCode}</li>
                                </ul>
                            </Disclosure.Panel>
                        </>
                    )}
                </Disclosure>
            </div>
        </Panel>
        <Panel className="flex flex-col md:w-1/3 sticky top-2">
            <div className="text-sm mt-4 p-4">
                <div className="flex justify-between">
                    <span>Totale </span> {Currency.defaultFormat(order.order.priceAmount, order.order.currency.iso_code)}
                </div>
                <div className="divider"></div>
                <div className="flex justify-between">
                    <span>Totale IVA (22%):</span> {Currency.defaultFormat(order.order.taxesAmount, order.order.currency.iso_code)}
                </div>
            </div>
            <div className="ml-auto mt-auto flex flex-col text-secondary text-sm">
                Totale con IVA
                <span className="text-4xl font-semibold text-primary">
                {Currency.defaultFormat(order.order.priceAmountTotal, order.order.currency.iso_code)}
            </span>
            </div>
            <Button buttonType="primary" className="w-full mt-4" onClick={ () => {
                mutate( {
                    variables: {
                        orderId: order.order.id,
                    }
                } ).then( () => {
                    navigate( `/order/details/${order.order.id}` );
                } );
            } }>
                Conferma ordine
            </Button>
        </Panel>
    </main>;
};

export default OrderCheckoutPage;