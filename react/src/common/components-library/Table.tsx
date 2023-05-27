import {faAngleLeft, faAngleRight} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useEffect} from "react";
import {ColumnDef, flexRender, getCoreRowModel, getPaginationRowModel, useReactTable} from "@tanstack/react-table";
import Button from "./Button";

type ReactTableProps<T extends object = any> = {
    data: T[]
    columns: ColumnDef<T>[]
    hideHeader?: boolean
    pageCount?: number | undefined
    updatePagination?: ( index: number, size: number ) => void
}
const Table: React.FC<ReactTableProps> = <T extends object>(props: ReactTableProps<T>) => {
    const table = useReactTable({
        data: props.data,
        columns: props.columns,
        getCoreRowModel: getCoreRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        pageCount: props.pageCount,
        manualPagination: true,
        initialState: {
            pagination: {
                pageIndex: 0
            }
        }
    }
    );

    useEffect( () => {
        if ( props.updatePagination ) {
            props.updatePagination( table.getState().pagination.pageIndex, table.getState().pagination.pageSize );
        }
    }, [table.getState().pagination.pageIndex] );

    return (
        <div className="flex flex-col">
            <div className="overflow-x-auto sm:-mx-6 lg:-mx-8">
                <div className="inline-block min-w-full py-4 sm:px-6 lg:px-8">
                    <div className="overflow-hidden p-2">
                        <table className="table w-full">
                            {
                                props.hideHeader ? '' :
                                <thead>
                                {table.getHeaderGroups().map((headerGroup) => (
                                    <tr key={headerGroup.id}>
                                        {headerGroup.headers.map((header) => (
                                            <th key={header.id}>
                                                {header.isPlaceholder ? null : flexRender(header.column.columnDef.header, header.getContext())}
                                            </th>
                                        ))}
                                    </tr>
                                ))}
                                </thead>
                            }
                            <tbody>
                            {table.getRowModel().rows.map((row) => (
                                <tr key={row.id}>
                                    {row.getVisibleCells().map((cell) => (
                                        <td key={cell.id}>
                                            {flexRender(cell.column.columnDef.cell, cell.getContext())}
                                        </td>
                                    ))}
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    </div>
                    <div className="w-full flex justify-center btn-group">
                            <Button type="secondary" onClick={table.previousPage} className={ table.getCanPreviousPage() ? '' : 'btn-disabled' }><FontAwesomeIcon icon={faAngleLeft}/></Button>
                            <span className="flex items-center p-2 uppercase text-sm font-semibold bg-secondary text-white text-center"> Pagina { table.getState().pagination.pageIndex + 1 } di { table.getPageCount() }</span>
                            <Button type="secondary" onClick={table.nextPage} className={ table.getCanNextPage() ? '' : 'btn-disabled' }><FontAwesomeIcon icon={faAngleRight}/></Button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Table;