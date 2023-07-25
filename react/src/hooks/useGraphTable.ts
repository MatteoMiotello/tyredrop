import { ColumnDef } from "@tanstack/react-table";
import { OperationVariables, QueryResult } from "@apollo/client";
import {useTable} from "../common/components/shelly-ui";
import {Pagination} from "../__generated__/graphql";

type UseGraphTableProps<T = any, TQuery = any, TVariables extends OperationVariables = OperationVariables> = {
	data: T[],
	query: QueryResult<TQuery, TVariables>
	columns: ColumnDef<T>[],
	paginator?: Pagination
}

export const useGraphTable = <T = any,  TQuery = any, TVariables extends OperationVariables = OperationVariables>( {data, query, columns, paginator}: UseGraphTableProps<T> ) => {
	return useTable( {
		data: data,
		onPaginationChange: (pageIndex, pageCount) => {
			console.log( pageIndex * pageCount );

			return query.refetch({
				pagination: {
					limit: paginator?.limit,
					offset: ( pageIndex - 1 ) * pageCount
				}
			});
		},
		columns: columns,
		pageCount: paginator?.pageCount as number,
		pageSize: paginator?.limit as number,
		currentPage: paginator?.currentPage as number
	} );
};