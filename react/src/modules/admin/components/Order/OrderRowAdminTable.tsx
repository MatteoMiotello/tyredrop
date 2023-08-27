import {QueryResult} from "@apollo/client";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {
    Order,
    OrderRow,
    OrderRowsQuery,
    OrderRowsQueryVariables,
    UpdateOrderRowMutation, UpdateOrderRowMutationVariables
} from "../../../../__generated__/graphql";
import {useMutation} from "../../../../common/backend/graph/hooks";
import {UPDATE_ORDER_ROW} from "../../../../common/backend/graph/mutation/order";
import {
    BasicTable, Button,
    Form,
    Input,
    Modal,
    TableButtons, useForm,
    useModal,
    useTable
} from "../../../../common/components/shelly-ui";
import {Currency} from "../../../../common/utilities/currency";

type OrderRowAdminTableProps = {
    query: QueryResult<OrderRowsQuery, OrderRowsQueryVariables>
    order: Order
}
const OrderRowAdminTable: React.FC<OrderRowAdminTableProps> = ({query, order}) => {
    const [rowToEdit, setRowToEdit] = useState<OrderRow | undefined>();
    const [mutate] = useMutation<UpdateOrderRowMutation, UpdateOrderRowMutationVariables>(UPDATE_ORDER_ROW);
    const columns: ColumnDef<OrderRow>[] = [
        {
            accessorKey: 'id',
            size: 10,
            header: "Numero",
            cell: (p) => <span> #{`${order.orderNumber}_${p.row.original?.id}`}</span>
        },
        {
            accessorKey: "productItem",
            header: "Prodotto",
            size: 250,
            cell: (props: CellContext<OrderRow, any>) =>
                <span className="text-xs"> {props.row.original.productItemPrice.productItem.product.name} </span>
        },
        {
            accessorKey: "productItem.product.brand.name",
            header: "Brand",
            cell: (props: CellContext<OrderRow, any>) =>
                <span> {props.row.original.productItemPrice.productItem.product.brand.name} </span>
        },
        {
            id: 'ean',
            header: "EAN",
            cell: (p) => <span
                className="font-semibold"> {p.row.original.productItemPrice.productItem.product.code} </span>
        },
        {
            id: 'supplier',
            header: "Fornitore",
            cell: (p) => <span
                className="font-semibold"> {p.row.original.productItemPrice.productItem.supplier.name} </span>
        },
        {
            accessorKey: "quantity",
            header: "Quantità",
            size: 10,
        },
        {
            accessorKey: "amount",
            header: "Prezzo",
            size: 250,
            cell: (p: CellContext<OrderRow, any>) => {
                if (!p.row.original.productItemPrice) {
                    return null;
                }

                return <div className="text-xs">
                    <p> Fornitore: {Currency.defaultFormat(p.row.original.productItemPrice.productItem.supplierPrice, p.row.original.productItemPrice.currency.iso_code)} </p>
                    <p> Unitario: {Currency.defaultFormat(p.row.original.productItemPrice.value, p.row.original.productItemPrice.currency.iso_code)} </p>
                    <p> Totale: {Currency.defaultFormat(p.getValue(), p.row.original.productItemPrice.currency.iso_code)} </p>
                </div>;
            }
        },
        {
            accessorKey: "additionsAmount",
            header: "PFU",
            size: 100,
            cell: (p: CellContext<OrderRow, any>) =>
                <span> {Currency.defaultFormat(p.getValue(), p.row.original.productItemPrice.currency.iso_code)} </span>
        },
        {
            accessorKey: "trackingNumber",
            header: "Tracking",
            size: 400,
            cell: (props) => <>
                {
                    props.row.original.trackingNumber &&
                    <a href={props.row.original.trackingNumber as string}
                       className="link-accent"> {props.row.original.trackingNumber} </a>
                }
            </>
        },
        {
            id: 'actions',
            header: "",
            size: 10,
            cell: (p: CellContext<OrderRow, any>) => <TableButtons>
                <TableButtons.Edit onClick={() => setRowToEdit(p.row.original)}/>
            </TableButtons>
        }
    ];

    const {table} = useTable<OrderRow>({
        columns: columns,
        data: query.data?.orderRows as OrderRow[],
    });

    const modal = useModal({
        onClose: () => setRowToEdit(undefined)
    });

    const form = useForm({
        onSuccess: () => {
            modal.close();
            query.refetch();
        }
    });

    useEffect(() => {
        if (!rowToEdit) {
            return;
        }
        form.setFormValues({
            trackingNumber: rowToEdit.trackingNumber
        });
        modal.open();

    }, [rowToEdit]);

    return <>
        <Modal modal={modal}>
            <Modal.Title>
                Modifica tracking
            </Modal.Title>
            <Form form={form} saveForm={(data) => {
                if (!rowToEdit) {
                    return false;
                }

                return mutate({
                    variables: {
                        rowId: rowToEdit.id,
                        input: {
                            trackingNumber: data.trackingNumber
                        }
                    }
                });
            }}>
                <Input.FormControl>
                    <Input.Label>
                        Tracking url
                    </Input.Label>
                    <Input {...form.registerInput({name: 'trackingNumber'})}></Input>
                </Input.FormControl>
                <Form.FormButtons>
                    <Button onClick={modal.close}>
                        Annulla
                    </Button>
                    <Button type="submit" buttonType="primary">
                        Salva
                    </Button>
                </Form.FormButtons>
            </Form>
        </Modal>
        <BasicTable table={table}/>
    </>;
};

export default OrderRowAdminTable;