import React, {ReactNode, memo} from "react";
import { Table, flexRender } from "@tanstack/react-table";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faArrowDown, faArrowUp} from "@fortawesome/free-solid-svg-icons";

type BasicTableProps<T = any> = {
    table: Table<T>
}

type SortingToggleProps = {
	order: "desc" | "asc" | false
}
const SortingToggle = memo<SortingToggleProps>( ({order}: SortingToggleProps) => {
	return <>
		{order === 'asc' && <FontAwesomeIcon icon={faArrowDown}/>}
		{order === 'desc' && <FontAwesomeIcon icon={faArrowUp}/>}
	</>;
} );

SortingToggle.displayName = 'SortingToggle';

const BasicTable: React.FC<BasicTableProps> = ( {table, ...props} ) => {
	return <div className="overflow-x-auto">
		<table className="table table-zebra">
			<thead>
				{table.getHeaderGroups().map(headerGroup => (
					<tr key={headerGroup.id}>
						{headerGroup.headers.map(header => {
							return (
								<th 
									key={header.id} 
									colSpan={header.colSpan}
									style={{
										width: header.getSize()
									}}
								>
									{header.isPlaceholder ? null : (
										<div className="flex justify-between">
											{
												flexRender(
													<button onClick={header.column.getToggleSortingHandler()} disabled={!header.column.columnDef.enableSorting}>{header.column.columnDef.header as ReactNode}</button>,
													header.getContext(),
											)}
											<SortingToggle order={header.column.getIsSorted()}/>
										</div>
									)}
								</th>
							);
						})}
					</tr>
				))}
			</thead>
			<tbody>
				{table.getRowModel().rows.map(row => {
					return (
						<tr key={row.id}>
							{row.getVisibleCells().map(cell => {
								return (
									<td 
										key={cell.id}
										style={{
											width: cell.column.getSize()
										}}
									>
										{flexRender(
											cell.column.columnDef.cell,
											cell.getContext()
										)}
									</td>
								);
							})}
						</tr>
					);
				})}
			</tbody>
		</table>
	</div>;
};

export default BasicTable;