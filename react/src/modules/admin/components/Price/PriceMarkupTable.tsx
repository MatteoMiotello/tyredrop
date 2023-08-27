import {QueryResult} from "@apollo/client";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React from "react";
import {PriceMarkupsQuery, PriceMarkupsQueryVariables, ProductPriceMarkup} from "../../../../__generated__/graphql";
import {BasicTable, TableButtons, useTable} from "../../../../common/components/shelly-ui";

type PriceMarkupTableProps = {
    query: QueryResult<PriceMarkupsQuery, PriceMarkupsQueryVariables>
    setMarkupToEdit: ( markup: ProductPriceMarkup | undefined ) => void
}
const PriceMarkupTable: React.FC<PriceMarkupTableProps> = ( {query, setMarkupToEdit} ) => {
    const columns: ColumnDef<ProductPriceMarkup>[] = [
        {
            id: 'element',
            header: 'Riferimento',
            cell: ( p ) => <div>
                {
                    p.row.original.brand && <span> Brand: { p.row.original.brand.name} </span>
                }
                {
                    p.row.original.productCategory && <span> Categoria: { p.row.original.productCategory.name} </span>
                }
                {
                    p.row.original.product && <span> Prodotto: { p.row.original.product.name} </span>
                }
                {
                    ( !p.row.original.brand &&
                    !p.row.original.productCategory &&
                    !p.row.original.product ) && <span> Default </span>
                }
            </div>
        },
        {
            accessorKey: 'markupPercentage',
            header: 'Percentuale',
            cell: (p: CellContext<ProductPriceMarkup, any>) => <span className="text-xl font-medium"> {p.getValue() as string}% </span>
        },
        {
            id: 'actions',
            cell: (p) => <TableButtons>
                <TableButtons.Edit onClick={()=> setMarkupToEdit( p.row.original )}/>
            </TableButtons>
        }
    ];

    const {table} = useTable<ProductPriceMarkup>({
        data: query.data?.priceMarkups as ProductPriceMarkup[],
        columns: columns
    });

    return <BasicTable table={table}/>;
};

export default PriceMarkupTable;