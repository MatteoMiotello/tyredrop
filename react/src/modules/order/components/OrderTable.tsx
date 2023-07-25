import React from "react";
import {Order, OrdersPaginator} from "../../../__generated__/graphql";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import {calculateTotal} from "../utils";
import {Link} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faMagnifyingGlass} from "@fortawesome/free-solid-svg-icons";
import OrderStatusBadge from "./OrderStatusBadge";
import Moment from "react-moment";
import {FilteredTable, Input, useForm} from "../../../common/components/shelly-ui";
import {useGraphTable} from "../../../hooks/useGraphTable";
import {QueryResult} from "@apollo/client";

type OrderTableProps = {
    query: QueryResult<{userOrders: OrdersPaginator}>
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
                    <li> {order.addressLine1}</li>
                    {order.addressLine2 && <li> {order.addressLine1}</li>}
                    <li> {`${order.city} (${order.province}) ${order.postalCode} ${order.country}`} </li>
                </ul>;
            }
        },
        {
            accessorKey: "status",
            header: "Stato",
            cell: ( props: CellContext<Order, any> ) => <OrderStatusBadge status={props.row.original.status}/>
        },
        {
            accessorKey: "orderRows",
            header: "Totale",
            cell: (props: CellContext<Order, any>) => <span> {calculateTotal(props.row.original)} </span>
        },
        {
            accessorKey: "id",
            id: "btn",
            header: "",
            size: 5,
            cell: (props: CellContext<Order, any>) => <Link to={'/order/details/' + props.getValue()}
                                                            className="btn btn-secondary"> <FontAwesomeIcon
                icon={faMagnifyingGlass}/> </Link>
        }
    ];

    const {table} = useGraphTable({
        data: props.query.data?.userOrders.data as Order[],
        columns: columns,
        query: props.query,
        paginator: props.query.data?.userOrders.pagination
    });
    const form = useForm();

    return <FilteredTable table={table} >
        <FilteredTable.FilterForm form={form} updateAsyncFilters={ (data) => console.log(  data ) }>
            <FilteredTable.FilterField>
                <Input.Label>
                    Data da
                </Input.Label>
                <Input type="date" {...form.registerInput({name:'date_from' }) }/>
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Data a
                </Input.Label>
                <Input type="date" {...form.registerInput({name:'date_to' }) }/>
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Numero
                </Input.Label>
                <Input type="text" {...form.registerInput({name:'number' }) }/>
            </FilteredTable.FilterField>
        </FilteredTable.FilterForm>

    </FilteredTable>;
};

export default OrderTable;