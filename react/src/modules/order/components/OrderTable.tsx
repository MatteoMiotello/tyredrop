import React from "react";
import Table from "../../../common/components-library/Table";
import {Order} from "../../../__generated__/graphql";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import {calculateTotal} from "../utils";
import {Link} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faMagnifyingGlass} from "@fortawesome/free-solid-svg-icons";

type OrderTableProps = {
    orders: Partial<Order>[]
}
const OrderTable: React.FC<OrderTableProps> = (props) => {
    const columns: ColumnDef<Order>[] = [
        {
            accessorKey: "id",
            header: "ID",
            cell: (props: CellContext<Order, any>) => <span>#{props.getValue()}</span>
        },
        {
            header: "Indirizzo",
            cell: ( props: CellContext<Order, any> ) => {
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
          header: "Stato"
        },
        {
            accessorKey: "orderRows",
            header: "Totale",
            cell: (props: CellContext<Order, any>) => <span> {calculateTotal(props.row.original)} </span>
        },
        {
            accessorKey: "id",
            header: "",
            size: 5,
            cell: ( props: CellContext<Order, any> ) => <Link to={'/order/details/'+props.getValue()} className="btn btn-secondary"> <FontAwesomeIcon icon={faMagnifyingGlass}/> </Link>
        }
    ];

    return <Table data={props.orders} columns={columns}></Table>;
};

export default OrderTable;