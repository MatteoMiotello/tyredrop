import {QueryResult} from "@apollo/client";
import {faCheck, faDownload, faEuro} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {ColumnDef} from "@tanstack/react-table";
import React from "react";
import Moment from "react-moment";
import {Link} from "react-router-dom";
import {AllInvoicesQuery, AllInvoicesQueryVariables, Invoice, InvoiceStatus} from "../../../../__generated__/graphql";
import {Button, FilteredTable, Input, Select, TableButtons, useForm} from "../../../../common/components/shelly-ui";
import {useGraphTable} from "../../../../hooks/useGraphTable";
import UserBillingSelect from "../User/UserBillingSelect";

type InvoiceAdminTableProps = {
    query: QueryResult<AllInvoicesQuery, AllInvoicesQueryVariables>,
    setInvoiceToDelete: (invoice: Invoice) => void
    setInvoiceToUpdate: (invoice: Invoice) => void
}

const InvoiceAdminTable: React.FC<InvoiceAdminTableProps> = ({query, setInvoiceToDelete, setInvoiceToUpdate}) => {
    const columns: ColumnDef<Invoice>[] = [
        {
            accessorKey: "number",
            header: "Numero",
        },
        {
            accessorKey: 'userBilling',
            header: "Utente",
            cell: (p) => <Link className="link-accent"
                               to={`/admin/user/${p.row.original.userBilling.user.id}`}> {p.row.original.userBilling.name} {p.row.original.userBilling.surname} </Link>
        },
        {
            accessorKey: 'createdAt',
            header: "Data",
            cell: (p) => <Moment date={p.getValue() as string} format="D/M/YYYY"></Moment>
        },
        {
            accessorKey: 'status',
            header: "Stato Pagamento",
            cell: (p) => <>
                {
                    p.getValue() == InvoiceStatus.ToPay ?
                        <span className="text-error font-semibold"> Da pagare </span>:
                        <span className="text-success font-semibold"> Pagata </span>
                }
            </>
        },
        {
            id: "actions",
            cell: (p) => <TableButtons>
                {
                    (p.row.original.status == InvoiceStatus.ToPay) ?
                        <div className="tooltip" data-tip="imposta come pagata">
                            <Button buttonType="success" size="sm" onClick={() => setInvoiceToUpdate(p.row.original)}>
                                <FontAwesomeIcon icon={faCheck}/>
                            </Button>
                        </div>
                        :
                        <div className="tooltip" data-tip="imposta come da pagare">
                            <Button buttonType="error" size="sm" onClick={() => setInvoiceToUpdate(p.row.original)}>
                                <FontAwesomeIcon icon={faEuro}/>
                            </Button>
                        </div>
                }

                <a href={p.row.original.fileUrl} download target="_blank" rel="noreferrer"
                   className="btn btn-sm btn-outline">
                    <FontAwesomeIcon icon={faDownload}/>
                </a>
                <TableButtons.Delete onClick={() => setInvoiceToDelete(p.row.original)}/>
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
            data: query.data?.allInvoices?.data as Invoice[],
            paginator: query.data?.allInvoices.pagination,
        }
    );

    return <FilteredTable table={table}>
        <FilteredTable.FilterForm form={form} updateAsyncFilters={(data) => query.refetch({
            input: {
                from: data.from,
                to: data.to,
                number: data.number,
                userBillingId: data.userBilling ?? null,
                status: data.status
            }
        })}>
            <FilteredTable.FilterField>
                <Input.Label>
                    Data da
                </Input.Label>
                <Input type="date" {...form.registerInput({name: 'from'})} />
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Data a
                </Input.Label>
                <Input type="date" {...form.registerInput({name: 'to'})} />
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Numero
                </Input.Label>
                <Input {...form.registerInput({name: 'number'})} />
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Pagamento
                </Input.Label>
                <Select {...form.registerInput({name: 'status'})} options={ [
                    {
                        value: InvoiceStatus.Payed,
                        title: "Pagate"
                    },
                    {
                        value: InvoiceStatus.ToPay,
                        title: "Da pagare"
                    }
                ] }/>
            </FilteredTable.FilterField>
            <Input.FormControl className="col-span-4">
                <Input.Label>
                    Intestatario
                </Input.Label>
                <UserBillingSelect small name="userBilling"/>
            </Input.FormControl>
        </FilteredTable.FilterForm>
    </FilteredTable>;
};
export default InvoiceAdminTable;