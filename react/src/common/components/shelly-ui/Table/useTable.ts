import {
	ColumnDef,
	ColumnSort,
	PaginationState,
	SortingState,
	getCoreRowModel,
	getFilteredRowModel,
	getPaginationRowModel, getSortedRowModel, useReactTable
} from "@tanstack/react-table";
import { useEffect, useState } from "react";

export type PaginationChangeHandler = (pageIndex: number, pageCount: number) => void
export type SortingChangeHandler = ( orders: { column: string, desc: boolean }[] ) => void

type UseTableProps<T = any> = {
	data: T[],
	columns: ColumnDef<T>[],
	onPaginationChange?: PaginationChangeHandler,
	onSortingChange?: SortingChangeHandler,
	pageSize?: number
	pageCount?: number
	defaultSorting?: ColumnSort[]
	currentPage?: number
}

const useTable = <T = any>({ data, columns, onPaginationChange, onSortingChange, defaultSorting, pageSize, pageCount, currentPage }: UseTableProps<T>) => {
	const [pagination, setPagination] = useState<PaginationState>({
		pageIndex: ( currentPage || 1 ) -1,
		pageSize: pageSize || 10,
	});

	const [sorting, setSorting] = useState<SortingState>(defaultSorting || []);

	useEffect( () => {
		if ( onPaginationChange ) {
			onPaginationChange(pagination.pageIndex + 1, pagination.pageSize );
		}
	}, [pagination]);

	useEffect( () => {
		if ( onSortingChange && sorting.length ) {
			onSortingChange( sorting.map( sort => ({column: sort.id, desc: sort.desc}) ) );
		}
	}, [sorting] );

	const table = useReactTable({
		data,
		columns,
		getCoreRowModel: getCoreRowModel(),
		getFilteredRowModel: getFilteredRowModel(),
		getPaginationRowModel: getPaginationRowModel(),
		getSortedRowModel: getSortedRowModel(),
		debugTable: true,
		onPaginationChange: setPagination,
		onSortingChange: setSorting,
		pageCount: pageCount,
		state: {
			pagination: pagination,
			sorting: sorting
		},
		manualPagination: true,
		manualFiltering: true,
		manualSorting: true
	});


	return { table };
};

export default useTable;