import {faCheck, faChevronUp} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Disclosure, RadioGroup} from "@headlessui/react";
import React, {useEffect, useState} from "react";
import {useLoaderData, useNavigate} from "react-router-dom";
import {
    AllMethodsQuery,
    AllMethodsQueryVariables,
    FetchOrderQuery,
    OrderStatus, PaymentMethod
} from "../../__generated__/graphql";
import Panel from "../../common/components-library/Panel";
import {Button} from "../../common/components/shelly-ui";
import {Currency} from "../../common/utilities/currency";
import OrderRowsTable from "./components/OrderRowsTable";
import {useMutation, useQuery} from "../../common/backend/graph/hooks";
import {ALL_PAYMENT_METHODS} from "../../common/backend/graph/query/payments";
import {PAY_ORDER} from "../../common/backend/graph/mutation/order";
import {useToast} from "../../store/toast";

const OrderCheckoutPage: React.FC = () => {
    const order = useLoaderData() as FetchOrderQuery;
    const navigate = useNavigate();
    const [selected, setSelected] = useState<PaymentMethod | undefined>();
    const {data} = useQuery<AllMethodsQuery, AllMethodsQueryVariables>(ALL_PAYMENT_METHODS);
    const [mutate] = useMutation(PAY_ORDER);
    const toastr = useToast();

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
                <Disclosure as="div" className="collapse">
                    {({open}) => (
                        <>
                            <Disclosure.Button
                                className="w-full flex justify-between collapse-title bg-base-200 items-center rounded-box">
                                <span>Pagamento</span>
                                <FontAwesomeIcon icon={faChevronUp} transform={{rotate: (open ? 0 : 180)}}/>
                            </Disclosure.Button>
                            <Disclosure.Panel className="m-4">
                                <RadioGroup value={selected} onChange={setSelected}>
                                    <RadioGroup.Label className="sr-only">Server size</RadioGroup.Label>
                                    <div className="space-y-2">
                                        {
                                            data?.paymentMethods &&
                                            data?.paymentMethods.map((method) => (
                                                <RadioGroup.Option
                                                    key={method?.name}
                                                    value={method}
                                                    className={({active, checked}) =>
                                                        `${
                                                            active
                                                                ? 'ring-2 ring-white ring-opacity-60 ring-offset-2 ring-offset-sky-300'
                                                                : ''
                                                        }
                  ${
                                                            checked ? 'bg-secondary bg-opacity-75 text-white' : 'bg-white'
                                                        }
                    relative flex cursor-pointer rounded-lg px-5 py-4 shadow-md focus:outline-none`
                                                    }
                                                >
                                                    {({active, checked}) => (
                                                        <>
                                                            <div className="flex w-full items-center justify-between">
                                                                <div className="flex items-center">
                                                                    <div className="text-sm">
                                                                        <RadioGroup.Label
                                                                            as="p"
                                                                            className={`font-medium  ${
                                                                                checked ? 'text-white' : 'text-gray-900'
                                                                            }`}
                                                                        >
                                                                            {method?.name}
                                                                        </RadioGroup.Label>
                                                                        <RadioGroup.Description
                                                                            as="span"
                                                                            className={`inline ${
                                                                                checked ? 'text-sky-100' : 'text-gray-500'
                                                                            }`}
                                                                        >
                                                                            <ul>
                                                                                <li>{method?.receiver}</li>
                                                                                <li> {method?.bank_name} </li>
                                                                                <li> {method?.iban} </li>
                                                                            </ul>
                                                                        </RadioGroup.Description>
                                                                    </div>
                                                                </div>
                                                                {checked && (
                                                                    <div className="shrink-0 text-white">
                                                                        <FontAwesomeIcon icon={faCheck}/>
                                                                    </div>
                                                                )}
                                                            </div>
                                                        </>
                                                    )}
                                                </RadioGroup.Option>
                                            ))}
                                    </div>
                                </RadioGroup>
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
                if ( !selected ) {
                    toastr.error( "Non è stato selezionato alcun metodo di pagamento." );
                    return;
                }

                mutate( {
                    variables: {
                        orderId: order.order.id,
                        methodCode: selected.code,
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