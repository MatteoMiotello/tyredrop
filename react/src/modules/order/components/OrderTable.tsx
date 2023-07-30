import React from "react";
import {FetchOrdersQuery, FetchOrdersQueryVariables, Order} from "../../../__generated__/graphql";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import {Link} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faHeadset, faMagnifyingGlass} from "@fortawesome/free-solid-svg-icons";
import OrderStatusBadge from "./OrderStatusBadge";
import Moment from "react-moment";
import {Button, FilteredTable, Input, TableButtons, useForm, useModal} from "../../../common/components/shelly-ui";
import {useGraphTable} from "../../../hooks/useGraphTable";
import {QueryResult} from "@apollo/client";
import {Currency} from "../../../common/utilities/currency";
import OrderHelpModal from "./OrderHelpModal";

type OrderTableProps = {
    query: QueryResult<FetchOrdersQuery, FetchOrdersQueryVariables>
}

const OrderTable: React.FC<OrderTableProps> = (props) => {

    const columns: ColumnDef<Order>[] = [
        {
            accessorKey: "id",
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
            accessorKey: "orderRows",
            header: "Totale",
            cell: (props: CellContext<Order, any>) => <span> {Currency.defaultFormat(
                props.row.original.priceAmount,
                props.row.original.currency.iso_code
            )} </span>
        },
        {
            id: "items",
            header: "Pezzi",
            cell: props => <ul className="list-disc text-sm">
                {
                    props.row.original.orderRows.map((row, key) => <li key={key} className="text-xs">
                        #{String(row?.id).padStart(6, '0')} {row && Currency.defaultFormat(row.amount, props.row.original.currency.iso_code)}
                    </li>)
                }
            </ul>
        },
        {
            accessorKey: "id",
            id: "btn",
            header: "",
            size: 5,
            cell: (props: CellContext<Order, any>) => <TableButtons>
                <Link to={'/order/details/' + props.getValue()} className="btn btn-secondary"> <FontAwesomeIcon
                    icon={faMagnifyingGlass}/> </Link>
                <Button buttonType="warning" onClick={modal.open}> <FontAwesomeIcon icon={faHeadset}/> </Button>
            </TableButtons>
        }
    ];

    const modal = useModal();

    const {table} = useGraphTable<Order, FetchOrdersQuery, FetchOrdersQueryVariables>({
        data: props.query.data?.userOrders?.data as Order[],
        columns: columns,
        query: props.query,
        paginator: props.query.data?.userOrders?.pagination,
    });
    const form = useForm({type: 'filter'});

    return <>
        <OrderHelpModal modal={modal}/>
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