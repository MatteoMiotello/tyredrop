import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {SearchQuery} from "../../../__generated__/graphql";
import Table from "../../../common/components-library/Table";
import TyreItemRow, {ProductRowItemData} from "./TyreItemRow";


type ProductTableProps = {
    products: SearchQuery
}

const ProductTable: React.FC<ProductTableProps> = (props) => {
    const [data, setData] = useState<(ProductRowItemData | null)[]>([]);
    const colums: ColumnDef<ProductRowItemData>[] = [
        {
            accessorKey: "content",
            cell: (props: CellContext<ProductRowItemData, any>) => <TyreItemRow data={props.row.original}/>
        }
    ];

    useEffect(() => {
        if (!props.products || !props.products.productItems || !props.products.productItems.productItems) {
            return;
        }

        const data = props.products.productItems.productItems.map((product): ProductRowItemData | null=> {
            if (!product || !product.price[0]) {
                return null;
            }

            return {
                brand: {
                    name: product.product.brand.name,
                    code: product.product.brand.code
                },
                name: product.product.name as string,
                code: product.product.code,
                price: {
                    value: product.price[0]?.value as number,
                    symbol: product.price[0]?.currency.symbol as string
                }
            };
        });

        if ( data !== undefined ) {
            setData(data);
        }
    }, [props.products]);

    return <Table
        hideHeader={true}
        data={data}
        columns={colums}
    />;
};

export default ProductTable;