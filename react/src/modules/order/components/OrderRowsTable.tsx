import React from "react";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import Table from "../../../common/components-library/Table";
import {FetchOrderQuery, OrderRow} from "../../../__generated__/graphql";
import {Currency} from "../../../common/utilities/currency";
import {Link} from "react-router-dom";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faMagnifyingGlass} from "@fortawesome/free-solid-svg-icons";

type OrderRowsTableProps = {
    order: FetchOrderQuery
}
const OrderRowsTable: React.FC<OrderRowsTableProps> = (props) => {
    const columns: ColumnDef<OrderRow>[] = [
        {
            accessorKey: "productItem",
            header: "Prodotto",
            cell: (props: CellContext<OrderRow, any>) => <span> {props.row.original.productItem.product.name} </span>
        },
        {
            accessorKey: "quantity",
            header: "Quantità",
        },
        {
            accessorKey: "amount",
            header: "Prezzo",
            cell: (p: CellContext<OrderRow, any>) => <span> {Currency.defaultFormat( p.getValue(), props.order.order.currency.iso_code )} </span>
        },
        {
            accessorKey: "productItemID",
            header: "",
            size: 5,
            cell: ( p: CellContext<OrderRow, any> ) => <Link to={"/products/details/" + p.row.original.productItem.id} className="btn btn-secondary"> <FontAwesomeIcon icon={faMagnifyingGlass}/> </Link>
        },
    ];

    return <Table data={props.order.order.orderRows} columns={columns} hidePagination={true}/>;
};

export default OrderRowsTable;