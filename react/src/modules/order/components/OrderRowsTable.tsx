import React from "react";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import {FetchOrderQuery, OrderRow} from "../../../__generated__/graphql";
import {Currency} from "../../../common/utilities/currency";
import {BasicTable, useTable} from "../../../common/components/shelly-ui";

type OrderRowsTableProps = {
    order: FetchOrderQuery
}
const OrderRowsTable: React.FC<OrderRowsTableProps> = (props) => {
    const columns: ColumnDef<OrderRow>[] = [
        {
            accessorKey: 'id',
            size: 10,
            header: "Numero",
            cell: (props) => <span> #{props.row.original.id} </span>
        },
        {
            accessorKey: "productItem",
            header: "Prodotto",
            cell: (props: CellContext<OrderRow, any>) => <span> {props.row.original.productItemPrice.productItem.product.name} </span>
        },
        {
            accessorKey: "quantity",
            header: "Quantità",
            size: 10,
        },
        {
            id: "unit_price",
            header: "Prezzo Unitario",
            cell: (p: CellContext<OrderRow, any>) => {
                if ( !p.row.original.productItemPrice ) {
                    return null;
                }

                return <span> {Currency.defaultFormat(p.row.original.productItemPrice.value, props.order.order.currency.iso_code)} </span>;
            }
        },
        {
            accessorKey: "amount",
            header: "Prezzo",
            cell: (p: CellContext<OrderRow, any>) => <span> {Currency.defaultFormat( p.getValue(), props.order.order.currency.iso_code )} </span>
        }
    ];

    const {table} = useTable<OrderRow>( {
        data: props.order.order.orderRows as OrderRow[],
        columns: columns
    } );

    return <BasicTable table={table}/>;
};

export default OrderRowsTable;