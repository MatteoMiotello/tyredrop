import React from "react";
import {QueryResult} from "@apollo/client";
import {AllOrdersQuery, AllOrdersQueryVariables, Order} from "../../../../__generated__/graphql";
import {FilteredTable, Input, TableButtons, useForm} from "../../../../common/components/shelly-ui";
import {useGraphTable} from "../../../../hooks/useGraphTable";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import {Currency} from "../../../../common/utilities/currency";
import {Link} from "react-router-dom";
import Moment from "react-moment";
import OrderStatusBadge from "../../../order/components/OrderStatusBadge";
import OrderStatusSelect from "../../../order/components/OrderStatusSelect";

type OrderTableProps = {
    query: QueryResult<AllOrdersQuery, AllOrdersQueryVariables>
}

const OrderAdminTable: React.FC<OrderTableProps> = ({query}) => {
    const columns: ColumnDef<Order>[] = [
        {
            accessorKey: 'orderNumber',
            header: "ID",
            size: 10,
            cell: (props: CellContext<Order, any>) => <span> #{props.getValue()} </span>
        },
        {
            accessorKey: 'createdAt',
            id: 'created_at',
            enableSorting: true,
            header: "Data",
            cell: (props) => <Moment date={props.row.original.createdAt}/>
        },
        {
            accessorKey: 'userBilling',
            header: "Utente",
            cell: (props: CellContext<Order, any>) => <div className="flex flex-col">
                <div>{`${props.getValue().name} ${props.getValue().surname}`}</div>
                <Link className="link-accent"
                      to={`/admin/user/${props.getValue().user.id}`}>{props.getValue().user.email}</Link>
            </div>
        },
        {
            accessorKey: 'status',
            header: "Stato",
            cell: (props) => <OrderStatusBadge status={props.row.original.status}/>
        },
        {
            accessorKey: 'priceAmount',
            header: "Totale",
            cell: (props) =>
                <span> {Currency.defaultFormat(props.row.original.priceAmount, props.row.original.currency.iso_code)} </span>
        },
        {
            id: 'actions',
            header: "",
            size: 10,
            cell: ( props ) => <TableButtons>
                    <TableButtons.Info to={`/admin/order/${props.row.original.id}`}/>
                </TableButtons>
        }
    ];

    const {table} = useGraphTable<Order, AllOrdersQuery, AllOrdersQueryVariables>({
        data: query.data?.allOrders.data as Order[],
        paginator: query.data?.allOrders.pagination,
        query: query,
        columns: columns
    });

    const form = useForm({type: 'filter'});

    return <FilteredTable table={table}>
        <FilteredTable.FilterForm form={form} updateAsyncFilters={(data) => {
            return query.refetch({
                filter: {
                    ...data,
                    status: data.status || null
                }
            });
        }}>
            <FilteredTable.FilterField>
                <Input.Label>
                    Numero
                </Input.Label>
                <Input {...form.registerInput({name: 'number'})}/>
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Data da
                </Input.Label>
                <Input type='date' {...form.registerInput({name: 'from'})}/>
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Data a
                </Input.Label>
                <Input type='date' {...form.registerInput({name: 'to'})}/>
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Stato
                </Input.Label>
                <OrderStatusSelect {...form.registerInput({name: 'status'})}/>
            </FilteredTable.FilterField>
        </FilteredTable.FilterForm>
    </FilteredTable>;
};

export default OrderAdminTable;