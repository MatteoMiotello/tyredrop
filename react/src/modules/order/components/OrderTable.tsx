import React, {useState} from "react";
import {FetchOrdersQuery, Order} from "../../../__generated__/graphql";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import {Link} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faHeadset, faMagnifyingGlass} from "@fortawesome/free-solid-svg-icons";
import admin from "../../../Admin";
import OrderStatusBadge from "./OrderStatusBadge";
import Moment from "react-moment";
import {Button, FilteredTable, Input, TableButtons, useForm, useModal} from "../../../common/components/shelly-ui";
import {useGraphTable} from "../../../hooks/useGraphTable";
import {QueryResult} from "@apollo/client";
import {Currency} from "../../../common/utilities/currency";
import OrderSupportModal from "./OrderSupportModal";

type OrderTableProps = {
    query: QueryResult<FetchOrdersQuery>
    admin?: boolean
}

const OrderTable: React.FC<OrderTableProps> = (props) => {
    const [selectedOrder, setSelectedOrder] = useState<Order | undefined>();
    const columns: ColumnDef<Order>[] = [
        {
            accessorKey: "orderNumber",
            size: 10,
            header: "ID",
            cell: (props: CellContext<Order, any>) => <span>#{props.getValue()}</span>
        },
        {
            accessorKey: "createdAt",
            id: "created_at",
            enableSorting: true,
            sortDescFirst: true,
            header: "Data",
            cell: props => {
                return <Moment date={props.row.original.createdAt}> </Moment>;
            }
        },
        {
            header: "Indirizzo",
            cell: (props: CellContext<Order, any>) => {
                const order = props.row.original;

                return <ul>
                    <li>{order.addressName}</li>
                    <li> {order.addressLine1}</li>
                    {order.addressLine2 && <li> {order.addressLine1}</li>}
                    <li> {`${order.city} (${order.province}) ${order.postalCode} ${order.country}`} </li>
                </ul>;
            }
        },
        {
            accessorKey: "status",
            header: "Stato",
            cell: (props: CellContext<Order, any>) => <OrderStatusBadge status={props.row.original.status}/>
        },
        {
            id: "total",
            header: "Totale",
            cell: (props: CellContext<Order, any>) => <span> {Currency.defaultFormat(
                props.row.original.priceAmount,
                props.row.original.currency.iso_code
            )} </span>
        },
        {
            id: "total_taxes",
            header: "Totale con iva",
            cell: (props: CellContext<Order, any>) => <span> {Currency.defaultFormat(
                props.row.original.priceAmountTotal,
                props.row.original.currency.iso_code
            )} </span>
        },
        {
            id: "items",
            header: "Pezzi",
            cell: props => <ul className="list-disc text-sm">
                {
                    props.row.original.orderRows.map((row, key) => <li key={key}
                                                                       className="text-xs flex justify-between">
                        <span
                            className="text-neutral-500">{`${ row && row.productItemPrice.productItem.product.code}`}</span> {row && row.quantity}
                    </li>)
                }
            </ul>
        },
        {
            accessorKey: "id",
            id: "btn",
            header: "",
            size: 5,
            cell: (p: CellContext<Order, any>) => <TableButtons>
                <Link to={props.admin ? `/admin/order/${p.getValue()}` : `/order/details/${p.getValue()}`}
                      className="btn btn-secondary"> <FontAwesomeIcon
                    icon={faMagnifyingGlass}/> </Link>
                {
                    !admin &&
                    <Button buttonType="warning" onClick={() => {
                        setSelectedOrder(p.row.original);
                        modal.open();
                    }}> <FontAwesomeIcon icon={faHeadset}/> </Button>
                }
            </TableButtons>
        }
    ];

    const modal = useModal();

    const {table} = useGraphTable<Order, FetchOrdersQuery>({
        data: props.query.data?.userOrders?.data as Order[],
        columns: columns,
        query: props.query,
        paginator: props.query.data?.userOrders?.pagination,
    });
    const form = useForm({type: 'filter'});

    return <>
        <OrderSupportModal modal={modal} order={selectedOrder}/>
        <FilteredTable table={table}>
            <FilteredTable.FilterForm form={form} updateAsyncFilters={(data) => {
                return props.query.refetch({
                    filter: data
                });
            }}>
                <FilteredTable.FilterField>
                    <Input.Label>
                        Data da
                    </Input.Label>
                    <Input type="date" {...form.registerInput({name: 'dateFrom'})}/>
                </FilteredTable.FilterField>
                <FilteredTable.FilterField>
                    <Input.Label>
                        Data a
                    </Input.Label>
                    <Input type="date" {...form.registerInput({name: 'dateTo'})}/>
                </FilteredTable.FilterField>
                <FilteredTable.FilterField>
                    <Input.Label>
                        Numero
                    </Input.Label>
                    <Input type="text" {...form.registerInput({name: 'number'})}/>
                </FilteredTable.FilterField>
            </FilteredTable.FilterForm>
        </FilteredTable>
    </>;
};

export default OrderTable;