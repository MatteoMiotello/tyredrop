import React from "react";
import Moment from "react-moment";
import {FilteredTable, Input, Select, TableButtons, useForm} from "../../../../common/components/shelly-ui";
import {useGraphTable} from "../../../../hooks/useGraphTable";
import {QueryResult} from "@apollo/client";
import {FetchAllUsersQuery, FetchAllUsersQueryVariables, User} from "../../../../__generated__/graphql";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faCircleCheck, faTimesCircle} from "@fortawesome/free-solid-svg-icons";

type UserAdminTableProps = {
    query: QueryResult<FetchAllUsersQuery, FetchAllUsersQueryVariables>
}
const UserAdminTable: React.FC<UserAdminTableProps> = ({query}) => {
    const columns: ColumnDef<User>[] = [
        {
            accessorKey: 'userCode',
            header: "ID",
            size: 10
        },
        {
            accessorKey: 'email',
            header: "Email"
        },
        {
            accessorKey: 'userRole.name',
            header: "Ruolo"
        },
        {
            accessorKey: 'name',
            header: "Nome",
            cell: (props: CellContext<User, any>) =>
                <span> {props.row.original.name} {props.row.original.surname} </span>
        },
        {
            accessorKey: "createdAt",
            header: "Data registrazione",
            cell: ( props ) => <Moment date={props.getValue() as string}/>
        },
        {
            accessorKey: 'confirmed',
            header: "Confermato",
            cell: (props: CellContext<User, any>) => <>
                {
                    props.row.original.confirmed
                    ? <FontAwesomeIcon icon={faCircleCheck} className="text-success"/>
                    : <FontAwesomeIcon icon={faTimesCircle} className="text-error"/>
                }
                {
                    props.row.original.rejected && <span className="text-error"> Rifiutato </span>
                }
            </>
        },
        {
            accessorKey: 'id',
            id: 'actions',
            header: "",
            cell: (props) => <TableButtons>
                <TableButtons.Info to={`/admin/user/${props.getValue()}`}/>
            </TableButtons>
        }
    ];

    const {table} = useGraphTable({
        query: query,
        data: query.data?.users?.data as User[],
        columns: columns,
        paginator: query.data?.users?.pagination
    });

    const form = useForm({type: 'filter'});

    return <FilteredTable table={table}>
        <FilteredTable.FilterForm form={form} updateAsyncFilters={(data) => {
            return query.refetch({
                filter: {
                    name: data.name,
                    email: data.email,
                    confirmed: (data.confirmed === undefined || data.confirmed === '' ) ? null : (data.confirmed == 1 ? true : false)
                },
            });
        }}>
            <FilteredTable.FilterField>
                <Input.Label>
                    Email
                </Input.Label>
                <Input {...form.registerInput({name: "email"})}/>
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Nome
                </Input.Label>
                <Input {...form.registerInput({name: "name"})}/>
            </FilteredTable.FilterField>
            <FilteredTable.FilterField>
                <Input.Label>
                    Confermato
                </Input.Label>
                <Select options={[
                    {
                        value: true,
                        title: 'Confermati'
                    },
                    {
                        value: false,
                        title: 'Non Confermati',
                    },
                ]} {...form.registerInput({name: "confirmed"})}/>
            </FilteredTable.FilterField>
        </FilteredTable.FilterForm>
    </FilteredTable>;
};

export default UserAdminTable;