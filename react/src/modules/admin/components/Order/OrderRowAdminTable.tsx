import {QueryResult} from "@apollo/client";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React from "react";
import {Order, OrderRow, OrderRowsQuery, OrderRowsQueryVariables} from "../../../../__generated__/graphql";
import {BasicTable, useTable} from "../../../../common/components/shelly-ui";
import {Currency} from "../../../../common/utilities/currency";

type OrderRowAdminTableProps = {
    query: QueryResult<OrderRowsQuery, OrderRowsQueryVariables>
    order: Order
}
const OrderRowAdminTable: React.FC<OrderRowAdminTableProps> = ({query, order}) => {
    const columns: ColumnDef<OrderRow>[] = [
        {
            accessorKey: 'id',
            size: 10,
            header: "Numero",
            cell: (p) => <span> #{`${order.orderNumber}_${p.row.original?.id}`}</span>
        },
        {
            accessorKey: "productItem",
            header: "Brand",
            size: 250,
            cell: (props: CellContext<OrderRow, any>) =>
                <span> {props.row.original.productItemPrice.productItem.product.name} </span>
        },
        {
            accessorKey: "productItem.product.brand.name",
            header: "Prodotto",
            cell: (props: CellContext<OrderRow, any>) => <span> {props.row.original.productItemPrice.productItem.product.brand.name} </span>
        },
        {
            id: 'ean',
            header: "EAN",
            cell: (p) => <span className="font-semibold"> {p.row.original.productItemPrice.productItem.product.code} </span>
        },
        {
            id: 'supplier',
            header: "Fornitore",
            cell: (p) => <span className="font-semibold"> {p.row.original.productItemPrice.productItem.supplier.name} </span>
        },
        {
            accessorKey: "quantity",
            header: "Quantità",
            size: 10,
        },
        {
            id: 'supplier_price',
            header: "Prezzo fornitore",
            cell: (p) => <span> {Currency.defaultFormat(p.row.original.productItemPrice.productItem.supplierPrice, p.row.original.productItemPrice.currency.iso_code)} </span>
        },
        {
            id: "unit_price",
            header: "Prezzo Unitario",
            size: 100,
            cell: (p: CellContext<OrderRow, any>) => {
                if (!p.row.original.productItemPrice) {
                    return null;
                }

                return <span> {Currency.defaultFormat(p.row.original.productItemPrice.value, p.row.original.productItemPrice.currency.iso_code)} </span>;
            }
        },
        {
            accessorKey: "amount",
            header: "Prezzo",
            size: 100,
            cell: (p: CellContext<OrderRow, any>) =>
                <span> {Currency.defaultFormat(p.getValue(), p.row.original.productItemPrice.currency.iso_code)} </span>
        },
        {
            accessorKey: "additionsAmount",
            header: "Tasse",
            size: 100,
            cell: (p: CellContext<OrderRow, any>) =>
                <span> {Currency.defaultFormat(p.getValue(), p.row.original.productItemPrice.currency.iso_code)} </span>
        }
    ];

    const {table} = useTable<OrderRow>({
        columns: columns,
        data: query.data?.orderRows as OrderRow[],
    });

    return <BasicTable table={table}/>;
};

export default OrderRowAdminTable;