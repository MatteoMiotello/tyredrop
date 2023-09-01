import {QueryResult} from "@apollo/client";
import {faDownload} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {ColumnDef} from "@tanstack/react-table";
import React from "react";
import Moment from "react-moment";
import {AllUserInvoicesQuery, AllUserInvoicesQueryVariables, Invoice} from "../../../__generated__/graphql";
import { FilteredTable, Input, TableButtons, useForm} from "../../../common/components/shelly-ui";
import {useGraphTable} from "../../../hooks/useGraphTable";

type InvoiceTableProps = {
    query: QueryResult<AllUserInvoicesQuery, AllUserInvoicesQueryVariables>
}

const InvoiceTable: React.FC<InvoiceTableProps> = ({query}) => {
    const columns: ColumnDef<Invoice>[] = [
        {
            accessorKey: "number",
            header: "Numero",
        },
        {
            accessorKey: 'createdAt',
            header: "Data",
            cell: (p) => <Moment date={p.getValue() as string} format="D/M/YYYY"></Moment>
        },
        {
            id: "actions",
            cell: (p) => <TableButtons>
                <TableButtons>
                    <a href={p.row.original.fileUrl} download target="_blank" rel="noreferrer" className="btn btn-sm btn-outline">
                        <FontAwesomeIcon icon={faDownload}/>
                    </a>
                </TableButtons>
            </TableButtons>
        }
    ];

    const form = useForm({
        type: 'filter'
    });
    const {table} = useGraphTable(
        {
            columns: columns,
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            query: query,
            data: query.data?.allInvoices.data as Invoice[],
            paginator: query.data?.allInvoices.pagination,
        }
    );

    return <FilteredTable table={table}>
        <FilteredTable.FilterForm form={form} updateAsyncFilters={(data) => query.refetch( {
            input: {
                from: data.from,
                to: data.to,
                number: data.number
            }
        } )}>
            <FilteredTable.FilterField>
                <Input.Label>
                    Data da
                </Input.Label>
                <Input type="date" {...form.registerInput({ name: 'from' })} />
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Data a
                </Input.Label>
                <Input type="date" {...form.registerInput({ name: 'to' })} />
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Numero
                </Input.Label>
                <Input {...form.registerInput({ name: 'number' })} />
            </FilteredTable.FilterField>
        </FilteredTable.FilterForm>
    </FilteredTable>;
};
export default InvoiceTable;